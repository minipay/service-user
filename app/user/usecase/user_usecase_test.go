package usecase_test

import (
	"cleanbase/app/user/mocks"
	"cleanbase/app/user/models"
	"cleanbase/app/user/usecase"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestStore(t *testing.T){
	mockUserRepo := new(mocks.Repository)
	mockUser := models.User{
		Id:   0,
		Name: "Nanda",
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.Id = 1
		mockUserRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

		u := usecase.NewUserUseCase(mockUserRepo, time.Second* 2)

		err := u.Store(context.TODO(), &tempMockUser)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
		mockUserRepo.AssertExpectations(t)
	})

}
