package config

import (
	"fmt"
	"os"
)

// Establish the app's configuration
// Should the app interface with an external database or services
// this would represent targets, ports, auth, etc by type
type config struct {
	Database *Database
}

type Database struct {
	ExercisesDataPath string
	UsersDataPath     string
	WorkoutsDataPath  string
}

// NewConfig returns the application's configuration object
func NewConfig() (*config, error) {
	ep := os.Getenv("EXERCISES_PATH")
	if ep == "" {
		return nil, fmt.Errorf("EXERCISES_PATH required")
	}

	up := os.Getenv("USERS_PATH")
	if up == "" {
		return nil, fmt.Errorf("USERS_PATH required")
	}

	wp := os.Getenv("WORKOUTS_PATH")
	if wp == "" {
		return nil, fmt.Errorf("WORKOUTS_PATH required")
	}

	cfg := &config{
		Database: &Database{
			ExercisesDataPath: ep,
			UsersDataPath:     up,
			WorkoutsDataPath:  wp,
		},
	}
	return cfg, nil
}
