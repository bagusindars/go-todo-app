package repositories

import (
	"database/sql"
	"errors"
	"simple-todo-app/internal/models"
)

type UserRepository interface {
	FindByEmail(email string) (models.Users, error)
	Create(data models.Users) (int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByEmail(email string) (models.Users, error) {
	var result models.Users

	err := r.db.QueryRow("SELECT id, name, email, password, created_at, updated_at from users WHERE email = $1", email).Scan(&result.Id, &result.Name, &result.Email, &result.Password, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Users{}, nil
		}
		return models.Users{}, errors.New("Error find user : " + err.Error())
	}

	return result, nil
}

func (r *userRepository) Create(data models.Users) (int, error) {
	var id int = 0

	err := r.db.QueryRow(
		"INSERT INTO users (name, email, password, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id",
		data.Name, data.Email, data.Password, data.CreatedAt, data.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return 0, errors.New("Error creating user : " + err.Error())
	}

	return id, nil
}
