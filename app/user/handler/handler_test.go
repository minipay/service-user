package handler_test

import (
	"cleanbase/app/user/handler"
	"cleanbase/app/user/mocks"
	"cleanbase/app/user/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	//"github.com/bxcodec/faker"
)

func TestStore(t *testing.T) {
	mockArticle := models.User{
		Id:     1,
		Name:   "Content",
	}

	tempMockArticle := mockArticle
	tempMockArticle.Id = 0
	mockUCase := new(mocks.Usecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := handler.UserHandler{
		ATUsecase: mockUCase,
	}
	err = handler.CreateUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

//func TestFetchs (t *testing.T){
//	var mockArticle models.User
//	err := faker.FakeData(&mockArticle)
//	assert.NoError(t, err)
//	mockUCase := new(mocks.Usecase)
//	mockListArticle := make([]*models.User, 0)
//	mockListArticle = append(mockListArticle, &models.User{})
//	num := 1
//	cursor := "2"
//	mockUCase.On("FetchUsers", mock.Anything, cursor, int64(num)).Return(mockListArticle, 10, nil)
//
//	e := echo.New()
//	req, err := http.NewRequest(echo.GET, "/alluser?num=1&cursor="+cursor, strings.NewReader(""))
//	assert.NoError(t, err)
//
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	hand := handler.UserHandler{
//		ATUsecase: mockUCase,
//	}
//	err = hand.FetchUsers(c)
//	require.NoError(t, err)
//
//	//responseCursor := rec.Header().Get("x-cursor")
//	//assert.Equal(t, "10", responseCursor)
//	assert.Equal(t, http.StatusOK, rec.Code)
//	mockUCase.AssertExpectations(t)
//}
