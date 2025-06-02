package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestAccountCreationAndLogin(t *testing.T) {
	testFile := "./test.db"
	db, err := sql.Open("sqlite3", testFile)
	if err != nil {
		t.Errorf("SQL database could not be opened")
	}
	defer func() {
		db.Close()
		os.Remove(testFile)
	}()
	DatabasePerformanceOptimisatioins(db)

	_, err = db.Exec(CreateTableQuery)
	if err != nil {
		t.Errorf("CreateTableQuery failed with %v", err)
	}

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
