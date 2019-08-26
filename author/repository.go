package author

import (
	"cleanbase/models"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetByID(c echo.Context, id int64) (models.Author, error)
}
