package service

import (
	"context"

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
	UpdateUserAccount(ctx context.Context, userid int64, params UserChangeInput) error
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
	return s.repo.List(ctx)
}

func (s service) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s service) RegisterUser(ctx context.Context, params UserChangeInput) (*model.User, error) {
	user, err := model.NewUser(params.Firstname, params.Lastname, params.Age)
	if err != nil {
		s.logger.Errorw("Invalid parameters were supplied to registration function", "params", params, "err", err)
		return nil, err
	}
	user, err = s.repo.Create(ctx, user)
	if err != nil {
		s.logger.Errorw("Error during the repository call", "error", err)
		return nil, err
	}
	return user, nil
}

func (s service) UpdateUserAccount(ctx context.Context, userid int64, params UserChangeInput) error {
	user, err := s.repo.GetByID(ctx, userid)
	if err != nil {
		s.logger.Errorw("Unable to fetch user for update process", "error", err)
		return err
	}

	user.Age = params.Age
	user.FirstName = params.Firstname
	user.LastName = params.Lastname

	if err := model.Validateuser(user); err != nil {
		s.logger.Errorw("New params for updated user failed validation", "params", params, "error", err)
		return err
	}
	err = s.repo.Update(ctx, user)
	if err != nil {
		s.logger.Errorw("Unable to update user in the repo", "error", err)
		return err
	}
	return nil
}

func (s service) SuspendUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
