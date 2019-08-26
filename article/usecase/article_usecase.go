package usecase

import (
	"cleanbase/article"
	"cleanbase/author"
	"cleanbase/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

type articleUsecase struct {
	articleRepo article.Repository
	authorRepo author.Repository
	contextTimeout time.Duration
}

func (a articleUsecase) Fetch(c echo.Context, num int64) (r []*models.Article, err error) {
	authorID := num

	data, err := a.articleRepo.Fecth(c, authorID)
	if err != nil {
		return nil, err
	}

	for i, item := range data {
		dataAuthor, er := a.authorRepo.GetByID(c, item.ID)
		if er != nil {
			logrus.Error(err)
			return nil, err
		}
		data[i].Author = dataAuthor
	}

	return data, nil
}

func NewArticleUseCase(a article.Repository, ar author.Repository, timeout time.Duration) article.Usecase{
	return &articleUsecase{
		articleRepo:a,
		authorRepo:ar,
		contextTimeout:timeout,
	}
}