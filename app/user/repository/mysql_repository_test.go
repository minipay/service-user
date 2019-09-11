package repository_test

import (
	"cleanbase/app/user/models"
	"cleanbase/app/user/repository"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestStore(t *testing.T) {
	u := &models.User{
		Id:   1,
		Name: "naufal",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT users SET name=\\?"
	prep := mock.ExpectPrepare(query)

	prep.ExpectExec().WithArgs(u.Name).WillReturnResult(sqlmock.NewResult(12, 1))

	a := repository.NewMysqlUserRepository(db)

	err = a.Store(context.TODO(), u)
	assert.NoError(t, err)
}
