package article

import (
	"cleanbase/models"
	"github.com/labstack/echo/v4"
)

type Usecase interface {
	Fetch(c echo.Context, num int64) ([]*models.Article, error)
}
