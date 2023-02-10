package handlers

import (
	"fmt"
	"time"
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

// GetUserYearlyExerciseTotals builds a map of years and months with summed
// reps and weight for a given user and exercise
func GetUserYearlyExerciseTotals(d dao, firstname, lastname, exerciseTitle string) (YearlyTotals, error) {
	yt := make(map[int]map[time.Month]Totals)
	user, err := findOneUser(d, models.NewUserQueryOptions(models.WithUserName(firstname, lastname)))
	if err != nil {
		return yt, err
	}

	exercise, err := findOneExercise(d, models.NewExerciseQueryOptions(models.WithExerciseTitle(exerciseTitle)))
	if err != nil {
		return yt, err
	}

	workouts := d.FindWorkouts(models.NewWorkoutQueryOptions(models.WithWorkoutUserId(user.Id), models.WithWorkoutExerciseId(exercise.Id)))

	for _, workout := range workouts {
		workoutYear := workout.DatetimeCompleted.Year()
		workoutMonth := workout.DatetimeCompleted.Month()
		blockTotals := blockTotals(workout.Blocks)
		// check if we already have an entry for the year of this workout
		if _, ok := yt[workoutYear]; !ok {
			// if not, add the entry
			yt[workoutYear] = map[time.Month]Totals{
				workoutMonth: blockTotals,
			}
		} else {
			// check if we already have an entry for the month of this workout
			if monthTotal, ok := yt[workoutYear][workoutMonth]; !ok {
				// if not, add the entry
				yt[workoutYear][workoutMonth] = blockTotals
			} else {
				// otherwise increment existing values
				monthTotal.totalReps += blockTotals.totalReps
				monthTotal.totalWeight += blockTotals.totalWeight
				yt[workoutYear][workoutMonth] = monthTotal
			}
		}
	}

	return yt, nil
}

// for convenience, GetMaxWeightMonths returns the "heaviest" month per year from YearlyTotals
func GetMaxWeightMonths(yt YearlyTotals) (map[int]string, error) {
	maxMonths := make(map[int]string)
	if len(yt) == 0 {
		return maxMonths, fmt.Errorf("unprocessable entity")
	}

	for year, months := range yt {
		var maxMonth string
		var maxWeight uint32
		for m, t := range months {
			if t.totalWeight > maxWeight {
				maxMonth = m.String()
				maxWeight = t.totalWeight
			}
		}
		maxMonths[year] = maxMonth
	}

	return maxMonths, nil
}

// GetTotalYearlyWeight sums weight from YearlyTotals
func GetTotalYearlyWeight(yt YearlyTotals) (map[int]uint32, error) {
	yearSummary := make(map[int]uint32)
	if len(yt) == 0 {
		return yearSummary, fmt.Errorf("unprocessable entity")
	}

	for year, months := range yt {
		var totalWeight uint32
		for _, t := range months {
			totalWeight += t.totalWeight
		}
		yearSummary[year] = totalWeight
	}

	return yearSummary, nil
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
