package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepository_FindByUserName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &UserRepository{db: db}

	userName := "JohnDoe"
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1", userName)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT \\* FROM users WHERE name = \\$1").
		WithArgs(userName).
		WillReturnRows(rows)
	mock.ExpectCommit()

	userNameObj, _ := NewUserName(userName)
	user, err := repo.FindByUserName(userNameObj)
	if err != nil {
		t.Errorf("FindByUserName() error = %v", err)
	}

	if user == nil {
		t.Errorf("FindByUserName() = nil, expected user object")
	} else if user.Name().value != userName {
		t.Errorf("FindByUserName() user name = %v, want %v", user.Name().value, userName)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Test case for when the user is not found
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT \\* FROM users WHERE name = \\$1").
		WithArgs("NonExistentUser").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	nonExistentUserName, _ := NewUserName("NonExistentUser")
	_, err = repo.FindByUserName(nonExistentUserName)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("FindByUserName() expected sql.ErrNoRows, got %v", err)
	}
}

func TestUserRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo, err := NewUserRepository(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new user repository", err)
	}

	user := User{
		id:   UserId{value: "1"},
		name: UserName{value: "JohnDoe"},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").WithArgs(user.id.value, user.name.value).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Save(&user)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
