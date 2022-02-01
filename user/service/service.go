package service

import (
	"context"
	"errors"

	"github.com/talkanbaev-artur/interview/user/model"
	"go.uber.org/zap"
)

type UserChangeInput struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Age       int    `json:"age"`
}

type Service interface {
	//queries
	ListUser(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)

	//mutations
	RegisterUser(ctx context.Context, params UserChangeInput) (*model.User, error)
	UpdateUserAccount(ctx context.Context, params UserChangeInput) error
	SuspendUser(ctx context.Context, id int64) error
}

type service struct {
	repo   Repository
	logger *zap.SugaredLogger
}

func NewService(r Repository, logger *zap.SugaredLogger) Service {
	return service{r, logger}
}

func (s service) ListUser(ctx context.Context) ([]*model.User, error) {
	return nil, errors.New("not implemented")
}

func (s service) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return nil, errors.New("not implemented")
}

func (s service) RegisterUser(ctx context.Context, params UserChangeInput) (*model.User, error) {
	return nil, errors.New("not implemented")
}

func (s service) UpdateUserAccount(ctx context.Context, params UserChangeInput) error {
	return errors.New("not implemented")
}

func (s service) SuspendUser(ctx context.Context, id int64) error {
	return errors.New("not implemented")
}
