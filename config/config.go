package config

// Establish the app's config
// Should the app interface with an external database or services
// this would represent targets, ports, auth
type config struct {
	exercisesDataPath string
	usersDataPath     string
	workoutsDataPath  string
}

func AppConfig() (*config, error) {
	// TODO read config kvs from environment
	return &config{
		exercisesDataPath: "dal/data/exercises.json",
		usersDataPath:     "dal/data/users.json",
		workoutsDataPath:  "dal/data/workouts.json",
	}, nil
}
