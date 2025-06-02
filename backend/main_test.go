package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestAccountCreationAndLogin(t *testing.T) {
	testFile := "./test.db"
	db := Unwrap(sql.Open("sqlite3", testFile))
	defer func() {
		db.Close()
		os.Remove(testFile)
	}()
	DatabasePerformanceOptimisatioins(db)

	Unwrap(db.Exec(CreateTableQuery))

	ShouldError(t, CallPasswordCheckQuery(
		db,
		"someone.something@somewhere.com",
		"password1234",
	))

	CallInsertUserQuery(
		db,
		"someone.something@somewhere.com",
		"password1234",
		"Spanish",
	)

	ShouldError(t, CallPasswordCheckQuery(
		db,
		"someone.something@somewhere.com",
		"password12345",
	))

	ShouldNotError(t, CallPasswordCheckQuery(
		db,
		"someone.something@somewhere.com",
		"password1234",
	))
}
