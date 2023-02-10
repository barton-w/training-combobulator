package handlers

import (
	"fmt"
	"time"
	"training-combobulator/models"
)

// dao defines the interface used by handlers
// the interface is implemented by the data access layer
type dao interface {
	FindUsers(opt models.UserQueryOptions) []models.User
	FindExercises(opt models.ExerciseQueryOptions) []models.Exercise
	FindWorkouts(opt models.WorkoutQueryOptions) []models.Workout
}

// YearlyTotals and Totals are data structures to aggregate
// workout data by year and month
type YearlyTotals map[int]map[time.Month]Totals

type Totals struct {
	totalReps   uint32
	totalWeight uint32
}

// convenience wrappers around doa for handler edge-cases and error handling
// this will return an error if 0 or duplicate users are found
func findOneUser(d dao, opt models.UserQueryOptions) (models.User, error) {
	users := d.FindUsers(opt)
	if len(users) != 1 {
		// could be refactored as a 404 in an http handler implementation
		return models.User{}, fmt.Errorf("user not found")
	}

	return users[0], nil
}

// this will return an error if 0 or duplicate exercises are found
func findOneExercise(d dao, opt models.ExerciseQueryOptions) (models.Exercise, error) {
	exercises := d.FindExercises(opt)
	if len(exercises) != 1 {
		// could be refactored as a 404 in an http handler implementation
		return models.Exercise{}, fmt.Errorf("exercise not found")
	}

	return exercises[0], nil
}
