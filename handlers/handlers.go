package handlers

import (
	"fmt"
	"training-combobulator/models"
)

// GetExerciseTotalWeight calculates the total weight lifted for a given exercise
func GetExerciseTotalWeight(d dao, exerciseTitle string) (uint32, error) {
	var sum uint32
	exercise, err := findOneExercise(d, models.NewExerciseQueryOptions(models.WithExerciseTitle(exerciseTitle)))
	if err != nil {
		return sum, err
	}

	workouts := d.FindWorkouts(models.NewWorkoutQueryOptions(models.WithWorkoutExerciseId(exercise.Id)))

	for _, workout := range workouts {
		blockTotal := blockTotals(workout.Blocks)
		sum += blockTotal.totalWeight
	}

	// if no workouts were found, this will return 0
	return sum, nil
}

// GetUserMonthlyExerciseTotals builds a map of month to summed reps and weight
// for a given user, exercise, and year
// Something could call this iteratively for a multi-year analysis if needed
func GetUserMonthlyExerciseTotals(d dao, firstname, lastname, exerciseTitle string, year int) (MonthlyTotals, error) {
	mt := make(MonthlyTotals, 12)
	user, err := findOneUser(d, models.NewUserQueryOptions(models.WithUserName(firstname, lastname)))
	if err != nil {
		return mt, err
	}

	exercise, err := findOneExercise(d, models.NewExerciseQueryOptions(models.WithExerciseTitle(exerciseTitle)))
	if err != nil {
		return mt, err
	}

	workouts := d.FindWorkouts(models.NewWorkoutQueryOptions(models.WithWorkoutUserId(user.Id), models.WithWorkoutExerciseId(exercise.Id)))

	for _, workout := range workouts {
		if workout.DatetimeCompleted.Year() == year {
			blockTotals := blockTotals(workout.Blocks)
			// check if we already have an entry for the month of this workout
			if monthTotal, ok := mt[workout.DatetimeCompleted.Month()]; !ok {
				// if not, add the entry to the map
				mt[workout.DatetimeCompleted.Month()] = blockTotals
			} else {
				// otherwise increment existing values
				monthTotal.totalReps += blockTotals.totalReps
				monthTotal.totalWeight += blockTotals.totalWeight
				mt[workout.DatetimeCompleted.Month()] = monthTotal
			}
		}
	}

	return mt, nil
}

// for convenience, GetMaxMonthWeight returns the "heaviest" month from a MonthlyTotals map as a string
func GetMaxMonthWeight(mt MonthlyTotals) (string, error) {
	var maxMonth string
	var maxWeight uint32
	if len(mt) == 0 {
		return maxMonth, fmt.Errorf("unprocessable entity")
	}

	for k, v := range mt {
		if v.totalWeight > maxWeight {
			maxMonth = k.String()
			maxWeight = v.totalWeight
		}
	}

	return maxMonth, nil
}

// GetTotalWeight sums weight from MonthlyTotals
func GetTotalWeight(mt MonthlyTotals) (uint32, error) {
	var totalWeight uint32
	if len(mt) == 0 {
		return totalWeight, fmt.Errorf("unprocessable entity")
	}

	for _, v := range mt {
		totalWeight += v.totalWeight
	}

	return totalWeight, nil
}

// GetUserPR finds a the pr for a given user and exercise
func GetUserPR(d dao, firstname, lastname, exerciseTitle string) (uint32, error) {
	var pr uint32
	user, err := findOneUser(d, models.NewUserQueryOptions(models.WithUserName(firstname, lastname)))
	if err != nil {
		return pr, err
	}

	exercise, err := findOneExercise(d, models.NewExerciseQueryOptions(models.WithExerciseTitle(exerciseTitle)))
	if err != nil {
		return pr, err
	}

	workouts := d.FindWorkouts(models.NewWorkoutQueryOptions(models.WithWorkoutUserId(user.Id), models.WithWorkoutExerciseId(exercise.Id)))

	for _, workout := range workouts {
		max := maxWeight(workout.Blocks)
		if max > pr {
			pr = max
		}
	}

	// if no workouts were found, this will return 0
	return pr, nil
}

// maxWeight returns the highest weight value from a workout block
func maxWeight(b []models.Block) uint32 {
	var max uint32
	for _, block := range b {
		for _, set := range block.Sets {
			if set.Weight != nil && *set.Weight > max {
				max = *set.Weight
			}
		}
	}
	return max
}

// blockTotals returns a Totals object from a workout block
func blockTotals(b []models.Block) Totals {
	t := Totals{}
	for _, block := range b {
		for _, set := range block.Sets {
			t.totalReps += set.Reps
			if set.Weight != nil {
				t.totalWeight += *set.Weight * set.Reps
			}
		}
	}
	return t
}
