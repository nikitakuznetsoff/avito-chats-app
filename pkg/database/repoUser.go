package database

import (
	"chatsapp/pkg/models"
)

func (repo *Repository) CreateUser(user *models.User) (int64, error){
	result, err := repo.DB.Exec(
		"INSERT INTO users (`username`) VALUES (?)",
		user.Username,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *Repository) GetUserByID(id uint32) (*models.User, error) {
	user := &models.User{}
	err := repo.DB.
		QueryRow("SELECT id, username, created_at " +
			"FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
