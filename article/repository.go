package article

import (
	"cleanbase/models"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Fecth(c echo.Context, num int64) (res []*models.Article, err error)
}
