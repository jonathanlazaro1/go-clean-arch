package repository

import (
	"database/sql"
	"log"

	"github.com/jonathanlazaro1/go-clean-arch/domain/model"
)

type userRepository struct {
	db *sql.DB
}

// UserRepository has contracts that match the UC interface of the same name, using PGSQL
type UserRepository interface {
	FindAll() ([]model.User, error)
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll() ([]model.User, error) {
	// create the select sql query
	sqlStatement := `SELECT id, name, age, created_at, updated_at, deleted_at FROM users`

	rows, err := ur.db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to query users. %f", err)
	}
	defer rows.Close()
	var users []model.User

	for rows.Next() {
		var user model.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

		if err != nil {
			log.Fatalf("Unable to scan row. %v", err)
		}
		users = append(users, user)
	}

	return users, err
}
