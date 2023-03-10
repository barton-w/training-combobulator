package models

import "time"

// Workout defines the data model for individual workout records
type Workout struct {
	UserId            uint32          `json:"user_id"`
	DatetimeCompleted WorkoutDateTime `json:"datetime_completed"`
	Blocks            []Block
}

// Block represents nested workout-block data
type Block struct {
	ExerciseId uint32 `json:"exercise_id"`
	Sets       []Set
}

// Set defines workout-set metadata
type Set struct {
	Reps   uint32  `json:"reps"`
	Weight *uint32 `json:"weight"`
}

// To conform to the datetime_completed format in the datasource
// we need a custom type wrapper and unmarshalling logic
type WorkoutDateTime struct {
	time.Time
}

func (t *WorkoutDateTime) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}

// WorkoutQueryOptions and its associated setting functions
// provide configurable query filters
// when interacting with the data access layer
type WorkoutQueryOptions struct {
	UserId     *uint32
	ExerciseId *uint32
}

type WorkoutOption func(*WorkoutQueryOptions)

func WithWorkoutUserId(id uint32) WorkoutOption {
	return func(wo *WorkoutQueryOptions) {
		wo.UserId = &id
	}
}

func WithWorkoutExerciseId(id uint32) WorkoutOption {
	return func(wo *WorkoutQueryOptions) {
		wo.ExerciseId = &id
	}
}

// NewWorkoutQueryOptions receives one or many option setting functions
func NewWorkoutQueryOptions(opts ...WorkoutOption) WorkoutQueryOptions {
	wo := &WorkoutQueryOptions{}
	for _, opt := range opts {
		opt(wo)
	}
	return *wo
}
