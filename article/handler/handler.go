package handler

import (
	"cleanbase/article"
	"cleanbase/author"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type PostGetAuthor struct {
	Author_id string   `json:"author_id" validate:"required"`
}

// ErroResponse ...
type ErroResponse struct {
	Message string `json:"message"`
}

type Data = interface{}
type Meta = interface{}

// ErroResponse ...
type SuccessResponse struct {
	Message string `json:"message"`
	Result Data `json:"result"`
	Meta Meta `json:"meta"`
}

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	AUsecase article.Usecase
	ATUsecase author.Usecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo, us article.Usecase, at author.Usecase) {
	handler := &ArticleHandler{
		AUsecase: us,
		ATUsecase: at,
	}
	e.GET("/articles", handler.FetchArticles)
	e.GET("/author", handler.FetchAuthor)
}

// FetchArticles ...
func (h ArticleHandler) FetchArticles(c echo.Context) (err error) {
	authorID, _ := strconv.ParseInt(c.QueryParam("author_id"), 10, 64)
	if err != nil {
		return err
	}

	listAr, err := h.AUsecase.Fetch(c, authorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErroResponse{Message: err.Error()})
	}

	res := SuccessResponse{
		Message:"a",
		Result:listAr,
	}
	return c.JSON(http.StatusOK, res)
}

// FetchAuthor ...
func (h ArticleHandler) FetchAuthor(c echo.Context) (err error) {
	authorID, _ := strconv.ParseInt(c.QueryParam("author_id"), 10, 64)

	listAr, err := h.ATUsecase.Fetch(c, authorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErroResponse{Message: err.Error()})
	}

	res := SuccessResponse{
		Message:"a",
		Result:listAr,
	}
	return c.JSON(http.StatusOK, res)
}