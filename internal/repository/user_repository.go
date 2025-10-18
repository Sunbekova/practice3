package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"practice4-sqlx/internal/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(user models.User) error {
	query := `
		INSERT INTO users (name, email, balance)
		VALUES (:name, :email, :balance)
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Select(&users, "SELECT id, name, email, balance FROM users ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT id, name, email, balance FROM users WHERE id = $1", id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}
