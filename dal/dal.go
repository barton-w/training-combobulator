package dal

import (
	"encoding/json"
	"os"
	"training-combobulator/config"
)

type dataAccess struct {
	Users []User
}

type User struct {
	Id        uint32 `json:"id"`
	FirstName string `json:"name_first"`
	LastName  string `json:"name_last"`
}

func Connect(cfg *config.Database) (*dataAccess, error) {
	usersBytes, err := os.ReadFile(cfg.UsersDataPath)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return nil, err
	}

	return &dataAccess{
		Users: users,
	}, nil
}
