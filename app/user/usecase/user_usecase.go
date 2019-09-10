package usecase

import (
	"cleanbase/app/user"
	"cleanbase/app/user/models"
	"context"
	"time"
)

type userUsecase struct {
	userRepo user.Repository
	contextTimeout time.Duration
}

func (h *userUsecase) FetchUsers(c context.Context, cursor string, num int64) (res []*models.User, nextCursor string,err error) {
	ctx, cancel := context.WithTimeout(c, h.contextTimeout)
	defer cancel()

	list, nextCursor, err := h.userRepo.FetchUsers(ctx, cursor, num)
	if err != nil {
		return nil, nextCursor, err
	}
	return list, nextCursor, nil
}

func (h *userUsecase) FetchUser(c context.Context, user int64) (res *models.User, err error) {
	ctx, cancel := context.WithTimeout(c, h.contextTimeout)
	defer cancel()

	res, err = h.userRepo.FetchUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *userUsecase) Store(c context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(c, h.contextTimeout)
	defer cancel()

	err := h.userRepo.Store(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func NewUserUseCase(ar user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:ar,
		contextTimeout:timeout,
	}
}