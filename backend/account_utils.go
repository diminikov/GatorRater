package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Account struct {
	ID       int64
	username string
	password string
}

// login queries for accounts that have the specified usernames and passwords.
func login(db *sql.DB, a Account) error {
	rows, err := db.Query("SELECT * FROM accounts WHERE username = ? AND password = ?;", a.username, a.password)
	if err != nil {
		return fmt.Errorf("Error %s while logging in", err)
	}
	defer rows.Close()

	return nil
}

func signup(db *sql.DB, a Account) error {
	query := "INSERT INTO accounts (username, password) VALUES (?, ?);"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}

	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, a.username, a.password)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d accounts created ", rows)
	return nil
}

func delete(db *sql.DB, a Account) error {
	rows, err := db.Query("DELETE FROM accounts WHERE username = ? AND password = ?;", a.username, a.password)
	if err != nil {
		return fmt.Errorf("Error %s while deleting", err)
	}
	defer rows.Close()

	return nil
}

func connect(username string, password string, database string) (*sql.DB, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   username,
		Passwd: password,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: database,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("Open %q", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("Ping %q", pingErr)
	}

	return db, nil
}
