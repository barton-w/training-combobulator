package handlers

import (
	"testing"
	"time"
	"training-combobulator/config"
	"training-combobulator/dal"
	"training-combobulator/models"
)

// For these tests I could also implement a fairly extensive mock
// of the dao interface, as opposed to useing the actual DAL with test data
// Gomock is a great package for that, but for simplicity I'm sticking to
// standard packages
func getTestDbConfig() *config.Database {
	return &config.Database{
		ExercisesDataPath: "../test/data/testExercises.json",
		UsersDataPath:     "../test/data/testUsers.json",
		WorkoutsDataPath:  "../test/data/testWorkouts.json",
	}
}

func TestGetExerciseTotalWeight(t *testing.T) {
	// Error
	dao := &mockDaoMultiReturn{}
	var expected uint32
	result, err := GetExerciseTotalWeight(dao, "someExercise")

	if err == nil {
		t.Fail()
		t.Logf("TestGetExerciseTotalWeight failed. expected error, got nil")
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetExerciseTotalWeight failed. expected: %v, got: %v", expected, result)
	}

	// Success
	da, err := dal.Connect(getTestDbConfig())
	expected = uint32(5400) //calculated from the test data
	result, err = GetExerciseTotalWeight(da, "Tractor Pulls")

	if err != nil {
		t.Fail()
		t.Logf("TestGetExerciseTotalWeight failed. unexpected error: %s", err.Error())
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetExerciseTotalWeight failed. expected: %v, got: %v", expected, result)
	}
}

func TestGetUserMonthlyExerciseTotals(t *testing.T) {
	// Errors
	dao := &mockDaoMultiReturn{}
	result, err := GetUserMonthlyExerciseTotals(dao, "first", "last", "someExercise", 2023)

	if err == nil {
		t.Fail()
		t.Logf("TestGetUserMonthlyExerciseTotals failed. expected error, got nil")
	}

	if len(result) > 0 {
		t.Fail()
		t.Logf("TestGetUserMonthlyExerciseTotals failed. unexpected return value")
	}

	// Success
	da, err := dal.Connect(getTestDbConfig())
	expected := MonthlyTotals{
		time.January:  Totals{totalReps: 9, totalWeight: 600},  //calculated from the test data
		time.February: Totals{totalReps: 15, totalWeight: 500}, //calculated from the test data
	}
	result, err = GetUserMonthlyExerciseTotals(da, "Alfred", "Music", "Bicep Curls", 2023)

	if err != nil {
		t.Fail()
		t.Logf("TestGetUserMonthlyExerciseTotals failed. unexpected error: %s", err.Error())
	}

	if result[time.January] != expected[time.January] && result[time.February] != expected[time.February] {
		t.Fail()
		t.Logf("TestGetExerciseTotalWeight failed. expected: %v, got: %v", expected, result)
	}
}

func TestGetMaxMonthWeight(t *testing.T) {
	// Error
	mt := MonthlyTotals{}
	expected := ""
	result, err := GetMaxMonthWeight(mt)

	if err == nil {
		t.Fail()
		t.Logf("TestGetMaxMonthWeight failed. expected error, got nil")
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetMaxMonthWeight failed. expected: %v, got: %v", expected, result)
	}

	mt[time.January] = Totals{totalWeight: 100}
	mt[time.February] = Totals{totalWeight: 200}
	mt[time.March] = Totals{totalWeight: 68}
	expected = "February"
	result, err = GetMaxMonthWeight(mt)

	if err != nil {
		t.Fail()
		t.Logf("TestGetMaxMonthWeight failed. unexpected error: %s", err.Error())
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetMaxMonthWeight failed. expected: %v, got: %v", expected, result)
	}
}

func TestGetTotalWeight(t *testing.T) {
	// Error
	mt := MonthlyTotals{}
	var expected uint32
	result, err := GetTotalWeight(mt)

	if err == nil {
		t.Fail()
		t.Logf("TestGetTotalWeight failed. expected error, got nil")
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetTotalWeight failed. expected: %v, got: %v", expected, result)
	}

	mt[time.January] = Totals{totalWeight: 100}
	mt[time.February] = Totals{totalWeight: 200}
	mt[time.March] = Totals{totalWeight: 68}
	expected = uint32(368)
	result, err = GetTotalWeight(mt)

	if err != nil {
		t.Fail()
		t.Logf("TestGetTotalWeight failed. unexpected error: %s", err.Error())
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetTotalWeight failed. expected: %v, got: %v", expected, result)
	}
}

func TestGetUserPR(t *testing.T) {
	// Error
	dao := &mockDaoMultiReturn{}
	var expected uint32
	result, err := GetUserPR(dao, "first", "last", "someExercise")

	if err == nil {
		t.Fail()
		t.Logf("TestGetUserPR failed. expected error, got nil")
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetUserPR failed. expected: %v, got: %v", expected, result)
	}

	// Success
	da, err := dal.Connect(getTestDbConfig())
	expected = uint32(1000) //calculated from the test data
	result, err = GetUserPR(da, "Train", "Heroic", "Tractor Pulls")

	if err != nil {
		t.Fail()
		t.Logf("TestGetUserPR failed. unexpected error: %s", err.Error())
	}

	if result != expected {
		t.Fail()
		t.Logf("TestGetUserPR failed. expected: %v, got: %v", expected, result)
	}
}

func TestMaxWeight(t *testing.T) {
	w1 := uint32(100)
	w2 := uint32(200)
	w3 := uint32(300)
	w4 := uint32(400)

	b := []models.Block{
		{
			Sets: []models.Set{
				{Weight: &w1}, {Weight: &w2}, {},
			},
		},
		{
			Sets: []models.Set{
				{Weight: &w4}, {Weight: &w3},
			},
		},
	}
	result := maxWeight(b)

	if result != w4 {
		t.Fail()
		t.Logf("TestMaxWeight failed. expected: %v, got: %v", w4, result)
	}
}

func TestBlockTotals(t *testing.T) {
	w1 := uint32(100)
	w2 := uint32(200)
	w3 := uint32(300)
	w4 := uint32(400)

	b := []models.Block{
		{
			Sets: []models.Set{
				{Reps: 1, Weight: &w1}, {Reps: 2, Weight: &w2}, {Reps: 3},
			},
		},
		{
			Sets: []models.Set{
				{Reps: 1, Weight: &w4}, {Reps: 2, Weight: &w3},
			},
		},
	}

	expected := Totals{
		totalReps:   9,
		totalWeight: 1500,
	}
	result := blockTotals(b)

	if result != expected {
		t.Fail()
		t.Logf("TestMaxWeight failed. expected: %v, got: %v", expected, result)
	}
}
