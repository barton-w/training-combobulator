package dal

import (
	"encoding/json"
	"os"
	"training-combobulator/config"
	"training-combobulator/models"
)

type dataAccess struct {
	exercises []models.Exercise
	users     []models.User
	workouts  []models.Workout
}

// Simulating connecting to a database, but instead of a client or connection
// being returned, I'm initializing data in memory and returning
// the implementation of an interface for data access
func Connect(cfg *config.Database) (*dataAccess, error) {
	exercisesBytes, err := os.ReadFile(cfg.ExercisesDataPath)
	if err != nil {
		return nil, err
	}

	var exercises []models.Exercise
	err = json.Unmarshal(exercisesBytes, &exercises)
	if err != nil {
		return nil, err
	}

	usersBytes, err := os.ReadFile(cfg.UsersDataPath)
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return nil, err
	}

	workoutsBytes, err := os.ReadFile(cfg.WorkoutsDataPath)
	if err != nil {
		return nil, err
	}

	var workouts []models.Workout
	err = json.Unmarshal(workoutsBytes, &workouts)
	if err != nil {
		return nil, err
	}

	return &dataAccess{
		exercises: exercises,
		users:     users,
		workouts:  workouts,
	}, nil
}

// dataAccess Find methods provide flexibility to handlers/consumers
// however the trade-off is the conditional options pattern used is not
// very extensible should models evolve
// Ideally we'd implement a data access interface via a real database client
func (da *dataAccess) FindUsers(opt models.UserQueryOptions) []models.User {
	users := make([]models.User, 0)

	switch {
	case opt.Id != nil:
		for _, user := range da.users {
			if user.Id == *opt.Id {
				users = append(users, user)
			}
		}
	case opt.Firstname != nil && opt.Lastname != nil:
		for _, user := range da.users {
			if user.Firstname == *opt.Firstname && user.Lastname == *opt.Lastname {
				users = append(users, user)
			}
		}
	}

	return users
}

func (da *dataAccess) FindExercises(opt models.ExerciseQueryOptions) []models.Exercise {
	exercises := make([]models.Exercise, 0)

	switch {
	case opt.Id != nil:
		for _, exercise := range da.exercises {
			if exercise.Id == *opt.Id {
				exercises = append(exercises, exercise)
			}
		}
	case opt.Title != nil:
		for _, exercise := range da.exercises {
			if exercise.Title == *opt.Title {
				exercises = append(exercises, exercise)
			}
		}
	}

	return exercises
}

func (da *dataAccess) FindWorkouts(opt models.WorkoutQueryOptions) []models.Workout {
	workouts := make([]models.Workout, 0)

	switch {
	case opt.UserId != nil && opt.ExerciseId != nil:
		for _, workout := range da.workouts {
			// here we need to return only blocks matching the requested ExerciseId
			if workout.UserId == *opt.UserId {
				blocks := filterBlocksByExerciseId(workout.Blocks, *opt.ExerciseId)
				workout.Blocks = blocks
				workouts = append(workouts, workout)
			}
		}
	case opt.UserId == nil && opt.ExerciseId != nil:
		for _, workout := range da.workouts {
			// here we need to return only blocks matching the requested ExerciseId
			blocks := filterBlocksByExerciseId(workout.Blocks, *opt.ExerciseId)
			workout.Blocks = blocks
			workouts = append(workouts, workout)
		}
	case opt.UserId != nil && opt.ExerciseId == nil:
		for _, workout := range da.workouts {
			if workout.UserId == *opt.UserId {
				workouts = append(workouts, workout)
			}
		}
	}

	return workouts
}

// convenience function for filtering workout-blocks
func filterBlocksByExerciseId(b []models.Block, exerciseId uint32) []models.Block {
	blocks := make([]models.Block, 0)

	for _, block := range b {
		if block.ExerciseId == exerciseId {
			blocks = append(blocks, block)
		}
	}

	return blocks
}
