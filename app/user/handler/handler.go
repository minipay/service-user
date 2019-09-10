package handler

import (
	"cleanbase/app/user"
	"cleanbase/app/user/models"
	"cleanbase/library"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	ATUsecase user.Usecase
}

func NewUserHandler(e *echo.Echo, at user.Usecase) {
	handler := &UserHandler{
		ATUsecase: at,
	}
	e.GET("/", handler.Home)
	e.POST("/user", handler.CreateUser)
	e.GET("/user", handler.FetchUser)
	e.GET("/alluser", handler.FetchUsers)
}

func (h *UserHandler) CreateUser(c echo.Context) (err error) {
	var post models.User
	err = c.Bind(&post)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := library.IsRequestValid(post); !ok && err != nil {
		return c.JSON(http.StatusBadRequest, library.ErroResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = h.ATUsecase.Store(ctx, &post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, library.ErroResponse{Message: err.Error()})
	}

	res := library.SuccessResponse{
		Message:"Success save",
		Result: post,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) FetchUser(c echo.Context) (err error){
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	idUser := c.QueryParam("user")

	if idUser == ""{
		return c.JSON(http.StatusBadRequest, library.ErroResponse{Message: "User tidak ditemukan"})
	}

	idU, err := strconv.ParseInt(c.QueryParam("user"), 10, 64)

	resUser, err := h.ATUsecase.FetchUser(ctx, idU)

	res := library.SuccessResponse{
		Message:"Success get data user",
		Result: resUser,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) FetchUsers(c echo.Context) (err error){
	numS := c.QueryParam("num")
	num, _ := strconv.ParseInt(numS, 10, 64)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	resUser, nextCursor, err := h.ATUsecase.FetchUsers(ctx, cursor, num)

	res := library.SuccessResponse{
		Message:"Success",
		Result: resUser,
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Home(c echo.Context) (err error){
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res := library.SuccessResponse{
		Message:"Success",
		Result: "Hello",
	}

	return c.JSON(http.StatusOK, res)
}