package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	// setup
	var err error
	db, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
	db.Exec("INSERT INTO users (id, name, active) VALUES (?, ?, ?)", "123", "user1", false)
	// run test
	retCode := m.Run()
	// tear down
	// exit
	os.Exit(retCode)
}

func TestDefaultIsolationLevel(t *testing.T) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()
	var isolationLevel string
	err = tx.QueryRow("SELECT @@TX_ISOLATION;").Scan(&isolationLevel)
	if err != nil {
		panic(err)
	}
	// if isolationLevel != sql.LevelRepeatableRead.String() {
	// 	t.Errorf("default isolation level is %s; want %s", isolationLevel, sql.LevelRepeatableRead)
	// }
	if isolationLevel != "REPEATABLE-READ" {
		t.Errorf("expect REPEATABLE-READ, but %s", isolationLevel)
	}
}

func TestRepetableRead(t *testing.T) {
	tx1, err := db.Begin()
	if err != nil {
		panic(err)
	}
	tx2, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx1.Commit()
	var active bool
	tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if active {
		t.Errorf("expect false, but %t", active)
	}
	tx2.Exec("UPDATE users SET active = ?", true)
	tx2.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if !active {
		t.Errorf("expect true, but %t", active)
	}
	tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if active {
		t.Errorf("expect false, but %t", active)
	}
	tx2.Commit()
	tx1.QueryRow("SELECT active FROM users WHERE id='123'").Scan(&active)
	if active {
		t.Errorf("expect false, but %t", active)
	}
	tx1.Exec("UPDATE users SET name = ?", "user2")
	tx1.Commit()
	var name string
	err = db.QueryRow("SELECT name, active FROM users WHERE id='123'").Scan(&name, &active)
	if err != nil {
		panic(err)
	}
	if !active {
		t.Errorf("expect false, but %t", active)
	}
	if name != "user2" {
		t.Errorf("expect user2, but %s", name)
	}
}
