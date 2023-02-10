package handlers

import (
	"testing"
	"training-combobulator/models"
)

type mockDaoMultiReturn struct{}

func (m *mockDaoMultiReturn) FindUsers(opt models.UserQueryOptions) []models.User {
	return []models.User{{Id: 1}, {Id: 2}}
}
func (m *mockDaoMultiReturn) FindExercises(opt models.ExerciseQueryOptions) []models.Exercise {
	return []models.Exercise{{Id: 3}, {Id: 4}}
}
func (m *mockDaoMultiReturn) FindWorkouts(opt models.WorkoutQueryOptions) []models.Workout {
	return []models.Workout{}
}

type mockDaoSingleReturn struct{}

func (m *mockDaoSingleReturn) FindUsers(opt models.UserQueryOptions) []models.User {
	return []models.User{{Id: 99, Firstname: "Wes", Lastname: "Barton"}}
}
func (m *mockDaoSingleReturn) FindExercises(opt models.ExerciseQueryOptions) []models.Exercise {
	return []models.Exercise{{Id: 100, Title: "Squat Thrusts"}}
}
func (m *mockDaoSingleReturn) FindWorkouts(opt models.WorkoutQueryOptions) []models.Workout {
	return []models.Workout{}
}

func TestFindOneUser(t *testing.T) {
	daoMulti := &mockDaoMultiReturn{}
	expected := models.User{}
	user, err := findOneUser(daoMulti, models.NewUserQueryOptions(models.WithUserId(1)))

	if err == nil {
		t.Fail()
		t.Logf("TestFindOneUser failed. expected error, got nil")
	}

	if user != expected {
		t.Fail()
		t.Logf("TestFindOneUser failed. expected: %v, got: %v", expected, user)
	}

	daoSingle := &mockDaoSingleReturn{}
	expected = models.User{Id: 99, Firstname: "Wes", Lastname: "Barton"}
	user, err = findOneUser(daoSingle, models.NewUserQueryOptions(models.WithUserId(1)))

	if err != nil {
		t.Fail()
		t.Logf("TestFindOneUser failed. unexpected error: %s", err.Error())
	}

	if user != expected {
		t.Fail()
		t.Logf("TestFindOneUser failed. expected: %v, got: %v", expected, user)
	}
}

func TestFindOneExercise(t *testing.T) {
	daoMulti := &mockDaoMultiReturn{}
	expected := models.Exercise{}
	exercise, err := findOneExercise(daoMulti, models.NewExerciseQueryOptions(models.WithExerciseId(1)))

	if err == nil {
		t.Fail()
		t.Logf("TestFindOneExercise failed. expected error, got nil")
	}

	if exercise != expected {
		t.Fail()
		t.Logf("TestFindOneExercise failed. expected: %v, got: %v", expected, exercise)
	}

	daoSingle := &mockDaoSingleReturn{}
	expected = models.Exercise{Id: 100, Title: "Squat Thrusts"}
	exercise, err = findOneExercise(daoSingle, models.NewExerciseQueryOptions(models.WithExerciseId(1)))

	if err != nil {
		t.Fail()
		t.Logf("TestFindOneExercise failed. unexpected error: %s", err.Error())
	}

	if exercise != expected {
		t.Fail()
		t.Logf("TestFindOneExercise failed. expected: %v, got: %v", expected, exercise)
	}
}
