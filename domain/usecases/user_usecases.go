package usecases

import (
	"context"

	"github.com/VitorFranciscoDev/trip-planner-api/domain/entities"
	"github.com/VitorFranciscoDev/trip-planner-api/domain/repositories"
)

type UserUseCases struct {
	r repositories.UserRepository
}

func NewUserUseCases(r repositories.UserRepository) *UserUseCases {
	return &UserUseCases{r: r}
}

func (u *UserUseCases) Login(ctx context.Context, email string, password string) (*entities.User, error) {
	user, err := u.r.Login(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCases) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := u.r.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCases) AddUser(ctx context.Context, user *entities.User) error {
	err := u.r.AddUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
