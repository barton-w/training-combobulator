package dal

import (
	"encoding/json"
	"os"
	"training-combobulator/config"
	"training-combobulator/models"
)

type database struct {
	users []models.User
}

func Connect(cfg *config.Database) (*database, error) {
	usersBytes, err := os.ReadFile(cfg.UsersDataPath)
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return nil, err
	}

	return &database{
		users: users,
	}, nil
}

func (db *database) FindUserByName(first, last string) *models.User {
	for _, user := range db.users {
		if user.FirstName == first && user.LastName == last {
			return &user
		}
	}
	return nil
}
