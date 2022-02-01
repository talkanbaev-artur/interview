package repo

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/talkanbaev-artur/interview/user/model"
	"github.com/talkanbaev-artur/interview/user/service"
)

type relUserRepository struct {
	r rel.Repository
}

func NewRelRepo(r rel.Repository) service.Repository {
	return relUserRepository{r}
}

func (r relUserRepository) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.r.FindAll(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r relUserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var u model.User
	err := r.r.Find(ctx, &u, rel.Where(where.Eq("id", id)))
	return &u, err
}

func (r relUserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.r.Insert(ctx, user)
	return user, err
}

func (r relUserRepository) Update(ctx context.Context, user *model.User) error {
	return r.r.Update(ctx, user)
}

func (r relUserRepository) Delete(ctx context.Context, userid int64) error {
	err := r.r.Transaction(ctx, func(ctx context.Context) error {
		cnt, err := r.r.DeleteAny(ctx, rel.From("users").Where(where.Eq("id", userid)))
		if err != nil {
			return err
		}
		if cnt != 1 {
			return service.ErrUserNotFound
		}
		return nil
	})
	return err
}
