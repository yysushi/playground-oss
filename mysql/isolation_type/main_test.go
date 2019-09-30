package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type MyTx struct {
	*sql.Tx
	done bool
}

func (tx *MyTx) Rollback() error{
	tx.done = true
	return tx.Tx.Rollback()
}

func (tx *MyTx) Commit() error{
	tx.done = true
	return tx.Tx.Commit()
}

var (
	db *sql.DB
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
	tx, err = db.Begin()
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
	var dropTable string = `
DROP TABLE IF EXISTS users
`
	_, err = db.Exec(dropTable)
	if err != nil {
		panic(err)
	}
	var createTable string = `
CREATE TABLE users (
  id varchar(255) NOT NULL,
  name text,
  active boolean,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name, active) VALUES (?, ?, ?)", "123", "user1", false)
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

func TestDefaultIsolationLevel(t *testing.T) {
	setUp()
	defer tearDown()
	var isolationLevel string
	err = tx1.QueryRow("SELECT @@TX_ISOLATION;").Scan(&isolationLevel)
	if err != nil {
		panic(err)
	}
	if isolationLevel != "REPEATABLE-READ" {
		t.Fatalf("expect REPEATABLE-READ, but %s", isolationLevel)
	}
}

func TestRepetableReadEvenOthersChange(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	_, err = tx2.Exec("UPDATE users SET active = ?", true)
	if err != nil {
		panic(err)
	}
	// when
	var active bool
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	// then
	if active {
		t.Fatalf("expect false, but %t", active)
	}
}

func TestRepetableReadEvenOthersCommitChange(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	_, err = tx2.Exec("UPDATE users SET active = ?", true)
	if err != nil {
		panic(err)
	}
	tx2.Commit()
	// when
	var active bool
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	// then
	if active {
		t.Fatalf("expect false, but %t", active)
	}
}
