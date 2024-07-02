package user

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserService_Exists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userService, err := NewUserService(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new user service", err)
	}

	userName := "JohnDoe"
	newUserName, _ := NewUserName(userName)
	user, _ := NewUser(*newUserName)

	query := "SELECT \\* FROM users WHERE username = \\$1"
	mock.ExpectBegin()
	// モックの結果セットに1行のデータを追加します。
	rows := sqlmock.NewRows([]string{"id", "username"}).AddRow("1", userName)
	mock.ExpectQuery(query).WithArgs(userName).WillReturnRows(rows)
	mock.ExpectCommit()

	exists, err := userService.Exists(user)
	if err != nil {
		t.Errorf("an error '%s' was not expected when checking if user exists", err)
	}

	if !exists {
		t.Errorf("expected user to exist")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
