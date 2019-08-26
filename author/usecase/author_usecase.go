package usecase

import (
	"cleanbase/author"
	"cleanbase/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

type authorUsecase struct {
	authorRepo author.Repository
	contextTimeout time.Duration
}

func (h authorUsecase) Fetch(c echo.Context, num int64) (models.Author, error) {
	data, err := h.authorRepo.GetByID(c, num)
	if err != nil {
		logrus.Error(err)
		//return nil, err
	}
	return data, nil
}

func NewAuthorUseCase(ar author.Repository, timeout time.Duration) author.Usecase{
	return &authorUsecase{
		authorRepo:ar,
		contextTimeout:timeout,
	}
}