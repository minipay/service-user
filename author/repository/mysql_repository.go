package repository

import (
	"cleanbase/author"
	"cleanbase/models"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
)

type mysqlAuthorRepository struct {
	DB *sql.DB
}

func (h mysqlAuthorRepository) GetByID(c echo.Context, id int64) (r models.Author, err error) {
	query := `SELECT id, name FROM authors WHERE id=?`
	row := h.DB.QueryRow(query, id)
	r = models.Author{}
	err = row.Scan(
		&r.ID,
		&r.Name,
	)
	if err != nil && err != sql.ErrNoRows {
		// log the error
		fmt.Println('a')
	}
	return
}

func NewMysqlAuthorRepository(db *sql.DB) author.Repository {
	return &mysqlAuthorRepository{
		DB:db,
	}
}