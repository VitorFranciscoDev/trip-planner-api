package repositories

import (
	"context"

	"github.com/VitorFranciscoDev/trip-planner-api/domain/entities"
)

type UserRepository interface {
	Login(ctx context.Context, email string, password string) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	AddUser(ctx context.Context, user *entities.User) error
}
