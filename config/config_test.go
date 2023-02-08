package config

import (
	"testing"
)

func TestAppConfig(t *testing.T) {
	expected := &config{
		exercisesDataPath: "dal/data/exercises.json",
		usersDataPath:     "dal/data/users.json",
		workoutsDataPath:  "dal/data/workouts.json",
	}
	cfg, err := AppConfig()
	if err != nil {
		t.Fatalf("AppConfig() failed. error: %s", err.Error())
	}
	if *cfg != *expected {
		t.Fatalf("AppConfig() failed. expected: %v, got: %v", expected, cfg)
	}
}
