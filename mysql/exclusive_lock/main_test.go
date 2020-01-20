package main

import (
	"context"
	"database/sql"
	"os"
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
	var dropAddressTable string = `
DROP TABLE IF EXISTS addresses
`
	_, err = db.Exec(dropAddressTable)
	if err != nil {
		panic(err)
	}
	var addressTable string = `
CREATE TABLE addresses (
  id varchar(255) NOT NULL,
  country_name text,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`
	_, err = db.Exec(addressTable)
	if err != nil {
		panic(err)
	}
	var userTable string = `
CREATE TABLE users (
  id varchar(255) NOT NULL,
  name text,
  address_id varchar(255) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (address_id) REFERENCES addresses (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`
	_, err = db.Exec(userTable)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO addresses (id, country_name) VALUES (?, ?)", "321", "japan")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO addresses (id, country_name) VALUES (?, ?)", "322", "usa")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name, address_id) VALUES (?, ?, ?)", "123", "user1", "321")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name, address_id) VALUES (?, ?, ?)", "124", "user2", "322")
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

// to same record
// exclusive select
// vs non exclusive select
func TestNotBlockNonExclusiveSelect(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	go func() {
		var any string
		err = tx2.QueryRow("SELECT id FROM users WHERE id='123'").Scan(&any)
		if err != nil {
			panic(err)
		}
		returned <- struct{}{}
	}()
	// when
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if blocked {
		t.Fatal("exclusive lock was blocked as unexpected")
	}
}

// to same record
// exclusive select
// vs exclusive select
func TestBlockExclusiveSelect(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	// when
	go func() {
		var any string
		err = tx2.QueryRowContext(ctx, "SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
		if err != context.Canceled {
			panic(err)
		}
		returned <- struct{}{}
	}()
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if !blocked {
		t.Fatal("exclusive lock was not blocked as unexpected")
	} else {
		cancel()
		<-returned
	}
}

// to same record
// exclusive select
// vs exclusive select
func TestRestartAfterCommit(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	go func() {
		var any string
		err = tx2.QueryRow("SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
		if err != nil {
			panic(err)
		}
		returned <- struct{}{}
	}()
	// when
	tx1.Commit()
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if blocked {
		t.Fatal("exclusive lock was blocked as unexpected")
	}
}

// to same record
// exclusive select
// vs update
func TestBlockUpdate(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	// when
	go func() {
		_, err = tx2.ExecContext(ctx, "UPDATE users SET name = 'masked' WHERE id='123'")
		if err != context.Canceled {
			panic(err)
		}
		returned <- struct{}{}
	}()
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if !blocked {
		t.Fatal("exclusive lock was not blocked as unexpected")
	} else {
		cancel()
		<-returned
	}
}

// to other record
// exclusive select (non-key)
// vs exclusive select (key)
func TestNonKeySelectBlockExclusiveSelect(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE name='user1' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	go func() {
		var any string
		err = tx2.QueryRow("SELECT id FROM users WHERE id='124' FOR UPDATE").Scan(&any)
		if err != nil {
			panic(err)
		}
		returned <- struct{}{}
	}()
	// when
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if !blocked {
		t.Fatal("exclusive lock was not blocked as unexpected")
	}
}

// to other record
// exclusive select (key)
// vs exclusive select (non-key)
func TestBlockExclusiveBigSelect(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	go func() {
		var any string
		err = tx2.QueryRow("SELECT id FROM users WHERE name='user2' FOR UPDATE").Scan(&any)
		if err != nil {
			panic(err)
		}
		returned <- struct{}{}
	}()
	// when
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if !blocked {
		t.Fatal("exclusive lock was not blocked as unexpected")
	}
}

// to other record
// exclusive select
// vs exclusive select
func TestNotBlockOtherExclusiveSelect(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE id='123' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	go func() {
		var any string
		err = tx2.QueryRow("SELECT id FROM users WHERE id='124' FOR UPDATE").Scan(&any)
		if err != nil {
			panic(err)
		}
		returned <- struct{}{}
	}()
	// when
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if blocked {
		t.Fatal("exclusive lock was blocked as unexpected")
	}
}

// to other record
// exclusive select (foreign key)
// vs exclusive select (key)
func TestForeignKeySelectNotBlockOtherKeyExclusiveSelect(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT id FROM users WHERE address_id='321' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	go func() {
		var any string
		err = tx2.QueryRow("SELECT id FROM users WHERE id='124' FOR UPDATE").Scan(&any)
		if err != nil {
			panic(err)
		}
		returned <- struct{}{}
	}()
	// when
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if blocked {
		t.Fatal("exclusive lock was blocked as unexpected")
	}
}

// to other record
// exclusive select (key)
// vs exclusive select (foreign key)
func TestNotBlockOtherForeignKeyExclusiveSelect(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	var any string
	err = tx1.QueryRow("SELECT name FROM users WHERE id='123' FOR UPDATE").Scan(&any)
	if err != nil {
		panic(err)
	}
	returned := make(chan struct{})
	go func() {
		var any string
		err = tx2.QueryRow("SELECT name FROM users WHERE address_id='322' FOR UPDATE").Scan(&any)
		if err != nil {
			panic(err)
		}
		returned <- struct{}{}
	}()
	// when
	var blocked bool
	select {
	case _ = <-returned:
		blocked = false
	case _ = <-time.After(100 * time.Millisecond):
		blocked = true
	}
	// then
	if blocked {
		t.Fatal("exclusive lock was blocked as unexpected")
	}
}
