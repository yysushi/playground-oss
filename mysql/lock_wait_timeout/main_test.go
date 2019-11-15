package main

import (
	"database/sql"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
	_, err = db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", "123", "user1")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", "124", "user2")
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

func TestLockWaitTimeout(t *testing.T) {
	setUp()
	defer tearDown()
	// given
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
	returned := make(chan struct{})
	errReturned := make(chan struct{})
	// when
	go func() {
		var any string
		rows, err = tx2.Query("SELECT id FROM users WHERE id='123' FOR UPDATE")
		if err != nil {
			if strings.Contains(err.Error(), "Error 1205") {
				errReturned <- struct{}{}
				return
			} else {
				panic(err)
			}
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&any); err != nil {
				if strings.Contains(err.Error(), "Error 1205") {
					errReturned <- struct{}{}
					return
				} else {
					panic(err)
				}
			}
		}
		if err := rows.Err(); err != nil {
			if strings.Contains(err.Error(), "Error 1205") {
				errReturned <- struct{}{}
				return
			} else {
				panic(err)
			}
		}
		returned <- struct{}{}
	}()
	select {
	case _ = <-errReturned:
		return
	case _ = <-returned:
		t.Fatal("exclusive lock was not blocked with unkonwn reason")
	case _ = <-time.After(5 * time.Second):
		t.Fatal("timeout")
	}
	// then
	tx1.Commit()
	<-returned
}
