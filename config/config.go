package config

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
	// TODO read config kvs from environment
	cfg := &config{
		Database: &Database{
			ExercisesDataPath: "dal/data/exercises.json",
			UsersDataPath:     "dal/data/users.json",
			WorkoutsDataPath:  "dal/data/workouts.json",
		},
	}
	return cfg, nil
}
