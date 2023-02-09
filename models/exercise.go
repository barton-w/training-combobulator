package models

// Exercise defines the data model for individual exercise records
type Exercise struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
}

// ExerciseQueryOptions and its associated setting functions
// provide configurable query filters
// when interacting with the data access layer
type ExerciseQueryOptions struct {
	Id    *uint32
	Title *string
}

type ExerciseOption func(*ExerciseQueryOptions)

func WithExerciseId(id uint32) ExerciseOption {
	return func(eo *ExerciseQueryOptions) {
		eo.Id = &id
	}
}

func WithExerciseTitle(title string) ExerciseOption {
	return func(eo *ExerciseQueryOptions) {
		eo.Title = &title
	}
}

// NewExerciseQueryOptions receives a given option-setting function
// This could be made variadic, looping over WithOption functions,
// should the data access layer support it
func NewExerciseQueryOptions(opt ExerciseOption) ExerciseQueryOptions {
	eo := &ExerciseQueryOptions{}
	opt(eo)
	return *eo
}
