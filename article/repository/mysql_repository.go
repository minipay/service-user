package repository

import (
	"cleanbase/article"
	"cleanbase/models"
	"database/sql"
	"github.com/labstack/echo/v4"

	"github.com/Masterminds/squirrel"
)

type mysqlArticleRepository struct {
	DB *sql.DB
}

func (h mysqlArticleRepository) Fecth(c echo.Context, authorID int64) (res []*models.Article, err error) {
	queryBuilder := squirrel.Select("id", "title", "content", "author_id").From("articles")
	//queryBuilder = queryBuilder.Limit(uint64(1))

	//if authorID != "" {
		queryBuilder = queryBuilder.Where(squirrel.Eq{
			"author_id": authorID,
		})
	//}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res = []*models.Article{}
	for rows.Next() {
		var item models.Article
		var authorID int64
		er := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Body,
			&authorID,
		)
		if er != nil {
			err = er
			return
		}

		item.Author = models.Author{ID: authorID}
		res = append(res, &item)
	}

	return
}

func NewMysqlArticleRepository(db *sql.DB) article.Repository{
	return &mysqlArticleRepository{
		DB: db,
	}
}