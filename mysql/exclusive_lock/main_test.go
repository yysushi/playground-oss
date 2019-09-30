package main

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db, db2 *sql.DB
)

func TestMain(m *testing.M) {
	// setup
	var err error
	db, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	db2, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	defer db2.Close()

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

func TestBlock(t *testing.T) {
	// given
	tx1, err1 := db.Begin()
	if err1 != nil {
		panic(err1)
	}
	tx2, err2 := db2.Begin()
	if err2 != nil {
		panic(err2)
	}
	defer tx2.Commit()
	var active bool
	back := make(chan struct{})
	err1 = tx1.QueryRow("SELECT active FROM users WHERE id='123' FOR UPDATE").Scan(&active)
	if err1 != nil {
		panic(err1)
	}
	// https://github.com/go-sql-driver/mysql/issues/731
	// ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()
	// when
	go func() {
		err2 = tx2.QueryRowContext(ctx, "SELECT active FROM users WHERE id='123' FOR UPDATE").Scan(&active)
		if err2 != nil {
			panic(err2)
		}
		back <- struct{}{}
	}()
	var backed bool
	select {
	case _ = <-back:
		backed = true
	case _ = <-time.After(100 * time.Millisecond):
		backed = false
	}
	// then
	if backed {
		t.Fatal("exclusive lock not work as expected")
	} else {
		// cancel()
		tx1.Commit()
		<-back
	}
}

// func TestNoBlockIfExclusiveLock(t *testing.T){
// 	tx1, err := db.Begin()
// 	if err != nil {
// 		panic(err)
// 	}
// 	tx2, err := db.Begin()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var active bool
// 	// tx1.QueryRow("SELECT active FROM users WHERE name='user1' FOR UPDATE").Scan(&active)
// 	tx1.QueryRow("SELECT active FROM users WHERE id='123' FOR UPDATE").Scan(&active)
// 	// tx2.QueryRow("SELECT active FROM users WHERE name='123' FOR UPDATE").Scan(&active)
// 	defer tx1.Commit()
// 	// tx2.QueryRow("SELECT 1").Scan(&active)
// 	// tx2.QueryRow("SELECT active FROM users WHERE name='123' FOR UPDATE").Scan(&active)
// 	// tx2.QueryRow("SELECT active FROM users WHERE name='123'").Scan(&active)
// 	// tx2.QueryRow("SELECT active FROM users WHERE name like '123' FOR UPDATE").Scan(&active)
// 	tx2.QueryRow("SELECT active FROM users WHERE name like '123' FOR UPDATE").Scan(&active)
// 	tx2.Commit()
// }

// func TestBlockIfExclusiveLock(t *testing.T){
// 	tx1, err := db.Begin()
// 	if err != nil {
// 		panic(err)
// 	}
// 	tx2, err := db.Begin()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var active bool
// 	tx1.QueryRow("SELECT active FROM users WHERE id='123' FOR UPDATE").Scan(&active)
// 	tx2.QueryRow("SELECT active FROM users WHERE id='123' FOR UPDATE").Scan(&active)
// 	tx1.Commit()
// 	tx2.Commit()
// }
