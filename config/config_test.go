package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// Errors
	os.Setenv("EXERCISES_PATH", "")
	cfg, err := NewConfig()

	if err == nil {
		t.Fail()
		t.Logf("TestNewConfig failed. error expected")
	}

	if cfg != nil {
		t.Fail()
		t.Logf("TestNewConfig failed. unexpected return: %v", cfg)
	}

	os.Setenv("EXERCISES_PATH", "something")
	os.Setenv("USERS_PATH", "")
	cfg, err = NewConfig()

	if err == nil {
		t.Fail()
		t.Logf("TestNewConfig failed. error expected")
	}

	if cfg != nil {
		t.Fail()
		t.Logf("TestNewConfig failed. unexpected return: %v", cfg)
	}

	os.Setenv("USERS_PATH", "something")
	os.Setenv("WORKOUTS_PATH", "")
	cfg, err = NewConfig()

	if err == nil {
		t.Fail()
		t.Logf("TestNewConfig failed. error expected")
	}

	if cfg != nil {
		t.Fail()
		t.Logf("TestNewConfig failed. unexpected return: %v", cfg)
	}

	// Success
	os.Setenv("EXERCISES_PATH", "dal/data/exercises.json")
	os.Setenv("USERS_PATH", "dal/data/users.json")
	os.Setenv("WORKOUTS_PATH", "dal/data/workouts.json")
	expected := &config{
		Database: &Database{
			ExercisesDataPath: "dal/data/exercises.json",
			UsersDataPath:     "dal/data/users.json",
			WorkoutsDataPath:  "dal/data/workouts.json",
		},
	}
	cfg, err = NewConfig()

	if err != nil {
		t.Fail()
		t.Logf("TestNewConfig failed. error: %s", err.Error())
	}

	if *cfg.Database != *expected.Database {
		t.Fail()
		t.Logf("TestNewConfig failed. expected: %v, got: %v", expected, cfg)
	}
}
