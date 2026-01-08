package repositories

import (
	"context"
	"database/sql"

	"github.com/VitorFranciscoDev/trip-planner-api/domain/entities"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) MySQLUserRepository {
	return MySQLUserRepository{db: db}
}

func (r MySQLUserRepository) Login(ctx context.Context, email string, password string) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE email = ? and password = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, email, password).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r MySQLUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE email = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r MySQLUserRepository) AddUser(ctx context.Context, user *entities.User) error {
	const query = `
		INSERT INTO users (uuid, name, email, password) VALUES (?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, user.UUID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
