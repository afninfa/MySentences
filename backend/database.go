package main

import (
	"database/sql"
	"net/mail"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var CreateTableQuery = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    targetLanguage TEXT NOT NULL
);
`

var InsertUserQuery = `
INSERT INTO users (email, password, targetLanguage) VALUES (?, ?, ?)
`

var PasswordCheckQuery = `
SELECT password FROM users WHERE email = ?
`

var CheckUserExistsQuery = `
SELECT COUNT(1) FROM users WHERE email = ?
`

func DatabasePerformanceOptimisatioins(db *sql.DB) {
	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA synchronous=NORMAL;")
	db.Exec("PRAGMA busy_timeout=5000;")
	db.Exec("PRAGMA wal_autocheckpoint=1000;")
}

func CallPasswordCheckQuery(db *sql.DB, email, password string) error {
	// valid email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrorEmailFormatting
	}
	var hashedPassword string
	err = db.QueryRow(PasswordCheckQuery, email).Scan(&hashedPassword)
	if err != nil {
		return ErrorUserNotFound
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	if err != nil {
		return ErrorWrongPassword
	}
	return nil
}

func CallInsertUserQuery(
	db *sql.DB,
	email,
	password,
	targetLanguage string,
) error {
	// valid email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrorEmailFormatting
	}

	// password hash
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	// email already used?
	var count int
	err = db.QueryRow(CheckUserExistsQuery, email).Scan(&count)
	if err != nil {
		return err
	} else if count > 0 {
		return ErrorEmailUsed
	}

	// execute query
	_, err = db.Exec(
		InsertUserQuery,
		email,
		string(hashedPassword),
		targetLanguage,
	)

	return err
}
