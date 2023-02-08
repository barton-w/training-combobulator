package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	expected := &config{
		Database: &Database{
			ExercisesDataPath: "dal/data/exercises.json",
			UsersDataPath:     "dal/data/users.json",
			WorkoutsDataPath:  "dal/data/workouts.json",
		},
	}
	cfg, err := NewConfig()
	if err != nil {
		t.Fatalf("AppConfig() failed. error: %s", err.Error())
	}
	if *cfg.Database != *expected.Database {
		t.Fatalf("AppConfig() failed. expected: %v, got: %v", expected, cfg)
	}
}
