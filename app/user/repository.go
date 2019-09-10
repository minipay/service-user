package user

import (
	"cleanbase/app/user/models"
	"context"
)

type Repository interface {
	Store(c context.Context, post *models.User) error
	FetchUser(c context.Context, user int64) (res *models.User, err error)
	FetchUsers(c context.Context, cursor string, num int64) (res []*models.User, nextCursor string, err error)
}