package main

import (
	"database/sql"
	"strings"
	"testing"

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
	db       *sql.DB
	tx1, tx2 *MyTx
	err      error
)

func setUp() {
	// db
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
	var dropTable2 string = `
DROP TABLE IF EXISTS users2
`
	_, err = db.Exec(dropTable2)
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
	_, err = db.Exec("INSERT INTO users (id, name, active) VALUES (?, ?, ?)", "124", "user2", false)
	if err != nil {
		panic(err)
	}
	var createTable2 string = `
CREATE TABLE users2 (
  id varchar(255) NOT NULL,
  name text,
  active boolean,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`
	_, err = db.Exec(createTable2)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users2 (id, name, active) VALUES (?, ?, ?)", "123", "user1", false)
	if err != nil {
		panic(err)
	}
	// transaction
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
	// transaction
	if !tx1.done {
		tx1.Commit()
	}
	if !tx2.done {
		tx2.Commit()
	}
	// db
	db.Close()
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

func TestNoDirtyRead(t *testing.T) {
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

// func TestRepeatableRead
func TestNoFuzzyRead(t *testing.T) {
	setUp()
	defer tearDown()
	// given1
	var active bool
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	if active {
		t.Fatalf("expect false, but %t", active)
	}
	// given2
	_, err = tx2.Exec("UPDATE users SET active = ?", true)
	if err != nil {
		panic(err)
	}
	tx2.Commit()
	// when
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	// then
	if active {
		t.Fatalf("expect false, but %t", active)
	}
}

// func TestRepeatableReadByOtherRowRead
func TestNoFuzzyReadByOtherRowRead(t *testing.T) {
	setUp()
	defer tearDown()
	// given1
	var active bool
	err = tx1.QueryRow("SELECT active FROM users WHERE id='124'").Scan(&active)
	if err != nil {
		panic(err)
	}
	if active {
		t.Fatalf("expect false, but %t", active)
	}
	// given2
	_, err = tx2.Exec("UPDATE users SET active = ?", true)
	if err != nil {
		panic(err)
	}
	tx2.Commit()
	// when
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	// then
	if active {
		t.Fatalf("expect false, but %t", active)
	}
}

func TestNoFuzzyReadByOtherTableRowRead(t *testing.T) {
	setUp()
	defer tearDown()
	// given1
	var active bool
	err = tx1.QueryRow("SELECT active FROM users2 WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	// given2
	_, err = tx2.Exec("UPDATE users SET active = ?", true)
	if err != nil {
		panic(err)
	}
	tx2.Commit()
	// when
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	// then
	if active {
		t.Fatalf("expect false, but %t", active)
	}
}

func TestSimilarDirtyReadWithLock(t *testing.T) {
	setUp()
	defer tearDown()
	// given
	_, err = tx2.Exec("UPDATE users SET active = ?", true)
	if err != nil {
		panic(err)
	}
	// when
	var active bool
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123' FOR UPDATE").Scan(&active)
	// then
	if strings.Index(err.Error(), "Error 1205: Lock wait timeout exceeded; try restarting transaction") == -1 {
		panic(err)
	}
}

func TestSimilarFuzzyReadWithLock(t *testing.T) {
	setUp()
	defer tearDown()
	// given1
	var active bool
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if err != nil {
		panic(err)
	}
	if active {
		t.Fatalf("expect false, but %t", active)
	}
	// given2
	_, err = tx2.Exec("UPDATE users SET active = ?", true)
	if err != nil {
		panic(err)
	}
	tx2.Commit()
	// when
	err = tx1.QueryRow("SELECT active FROM users WHERE id='123' FOR UPDATE").Scan(&active)
	if err != nil {
		panic(err)
	}
	// then
	if !active {
		t.Fatalf("expect true, but %t", active)
	}
}
