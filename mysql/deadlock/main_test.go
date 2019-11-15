package main

import (
	"database/sql"
	"os"
	"strings"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type MyError struct {
	idx  int
	orig error
}

func (e *MyError) Error() string {
	return e.orig.Error()
}

type MyTx struct {
	*sql.Tx
	done bool
}

func (tx *MyTx) Rollback() error {
	tx.done = true
	return tx.Tx.Rollback()
}

func (tx *MyTx) Commit() error {
	tx.done = true
	return tx.Tx.Commit()
}

var (
	db, db2  *sql.DB
	tx1, tx2 *MyTx
	err      error
)

func setUp() {
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		panic(err)
	}
	tx1 = &MyTx{tx, false}
	tx, err = db2.Begin()
	if err != nil {
		panic(err)
	}
	tx2 = &MyTx{tx, false}
}

func tearDown() {
	if !tx1.done {
		tx1.Commit()
	}
	if !tx2.done {
		tx2.Commit()
	}
}

func TestMain(m *testing.M) {
	// setup
	db, err = sql.Open("mysql", "root:password@/mydb")
	if err != nil {
		panic(err)
	}
	db2, err = sql.Open("mysql", "root:password@/mydb")
	if err != nil {
		panic(err)
	}
	var dropUserTable string = `
DROP TABLE IF EXISTS users
`
	_, err = db.Exec(dropUserTable)
	if err != nil {
		panic(err)
	}
	var userTable string = `
CREATE TABLE users (
  id varchar(255) NOT NULL,
  name text,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`
	_, err = db.Exec(userTable)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", "123", "userA")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", "124", "userB")
	if err != nil {
		panic(err)
	}
	// run test
	retCode := m.Run()
	// tear down
	db.Close()
	// exit
	os.Exit(retCode)
}

// func lockAB(tx1 *MyTx, doneLockA chan<- struct{}, crushedLockB <-chan struct{}) error {
func lockAB(tx1 *MyTx, doneLockA chan<- struct{}, doneLockB <-chan struct{}) error {
	var any string
	var rows *sql.Rows
	rows, err = tx1.Query("SELECT id FROM users WHERE id='123' FOR UPDATE")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&any); err != nil {
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	doneLockA <- struct{}{}
	<-doneLockB
	rows, err = tx1.Query("SELECT id FROM users WHERE id='124' FOR UPDATE")
	if err != nil {
		return &MyError{
			orig: err,
			idx:  1,
		}
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&any); err != nil {
			return &MyError{
				orig: err,
				idx:  2,
			}
		}
	}
	if err := rows.Err(); err != nil {
		return &MyError{
			orig: err,
			idx:  3,
		}
	}
	return nil
}

func lockBA(tx2 *MyTx, doneLockA <-chan struct{}, doneLockB chan<- struct{}) error {
	var any string
	var rows *sql.Rows
	rows, err = tx2.Query("SELECT id FROM users WHERE id='124' FOR UPDATE")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&any); err != nil {
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	<-doneLockA
	doneLockB <- struct{}{}
	rows, err = tx2.Query("SELECT id FROM users WHERE id='123' FOR UPDATE")
	if err != nil {
		return &MyError{
			orig: err,
			idx:  1,
		}
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&any); err != nil {
			return &MyError{
				orig: err,
				idx:  2,
			}
		}
	}
	if err := rows.Err(); err != nil {
		return &MyError{
			orig: err,
			idx:  3,
		}
	}
	return nil
}

func TestDeadLock(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var errs []error
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	doneLockA := make(chan struct{})
	doneLockB := make(chan struct{})
	wg.Add(1)
	go func() {
		err := lockAB(tx1, doneLockA, doneLockB)
		if err != nil {
			mutex.Lock()
			errs = append(errs, err)
			mutex.Unlock()
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		err := lockBA(tx2, doneLockA, doneLockB)
		if err != nil {
			mutex.Lock()
			errs = append(errs, err)
			mutex.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	if len(errs) != 1 {
		t.Fatal("num of errors should be one")
	}
	if !strings.Contains(errs[0].Error(), "Error 1213") {
		t.Fatal("error should be deadlock")
	}
	myErr, _ := errs[0].(*MyError)
	t.Logf("error idx is %d", myErr.idx)
}
