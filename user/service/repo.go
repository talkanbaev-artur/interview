package service

import (
	"context"
	"errors"

	"github.com/talkanbaev-artur/interview/user/model"
)

var (
	ErrUserNotFound = errors.New("user with such id not found")
)

type Repository interface {
	//queries
	List(ctx context.Context) ([]*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	//mutations
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error //dirty update
	Delete(ctx context.Context, userid int64) error
}
