package main

import (
	"fmt"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

const (
	Name = "ducai"
	Price = 8.89
)

func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("update user").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("insert into user").WithArgs(Name, Price).WillReturnResult(
		sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	if err := recordStats(db, Name, Price); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldRollbackStatUpdateOnFailture(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("update user").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("insert into user").
		WithArgs(Name, Price).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	if err = recordStats(db, Name, Price); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}