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

func TestGetUserYearlyExerciseTotals(t *testing.T) {
	// Errors
	dao := &mockDaoMultiReturn{}
	result, err := GetUserYearlyExerciseTotals(dao, "first", "last", "someExercise")

	if err == nil {
		t.Fail()
		t.Logf("TestGetUserYearlyExerciseTotals failed. expected error, got nil")
	}

	if len(result) > 0 {
		t.Fail()
		t.Logf("TestGetUserYearlyExerciseTotals failed. unexpected return value")
	}

	// Success
	da, err := dal.Connect(getTestDbConfig())
	expected := YearlyTotals{
		2023: map[time.Month]Totals{
			time.January:  {totalReps: 9, totalWeight: 600},  //calculated from the test data
			time.February: {totalReps: 15, totalWeight: 500}, //calculated from the test data
		},
	}
	result, err = GetUserYearlyExerciseTotals(da, "Alfred", "Music", "Bicep Curls")

	if err != nil {
		t.Fail()
		t.Logf("TestGetUserYearlyExerciseTotals failed. unexpected error: %s", err.Error())
	}

	if result[2023][time.January] != expected[2023][time.January] && result[2023][time.February] != expected[2023][time.February] {
		t.Fail()
		t.Logf("TestGetUserYearlyExerciseTotals failed. expected: %v, got: %v", expected, result)
	}
}

func TestGetMaxWeightMonths(t *testing.T) {
	// Error
	yt := YearlyTotals{}
	result, err := GetMaxWeightMonths(yt)

	if err == nil {
		t.Fail()
		t.Logf("GetMaxWeightMonths failed. expected error, got nil")
	}

	if len(result) > 0 {
		t.Fail()
		t.Logf("GetMaxWeightMonths failed. unexpected response: %v", result)
	}

	yt[2022] = map[time.Month]Totals{
		time.January:  {totalWeight: 100},
		time.February: {totalWeight: 200},
	}
	yt[2023] = map[time.Month]Totals{
		time.August:  {totalWeight: 999},
		time.October: {totalWeight: 68},
	}
	result, err = GetMaxWeightMonths(yt)

	if err != nil {
		t.Fail()
		t.Logf("GetMaxWeightMonths failed. unexpected error: %s", err.Error())
	}

	if result[2022] != "February" || result[2023] != "August" {
		t.Fail()
		t.Logf("GetMaxWeightMonths failed. incorrect data returned")
	}
}

func TestGetTotalYearlyWeight(t *testing.T) {
	// Error
	yt := YearlyTotals{}
	result, err := GetTotalYearlyWeight(yt)

	if err == nil {
		t.Fail()
		t.Logf("TestGetTotalYearlyWeight failed. expected error, got nil")
	}

	if len(result) > 0 {
		t.Fail()
		t.Logf("TestGetTotalYearlyWeight failed. unexpected response: %v", result)
	}

	yt[2022] = map[time.Month]Totals{
		time.January:  {totalWeight: 100},
		time.February: {totalWeight: 200},
	}
	yt[2023] = map[time.Month]Totals{
		time.August:  {totalWeight: 999},
		time.October: {totalWeight: 68},
	}
	result, err = GetTotalYearlyWeight(yt)

	if err != nil {
		t.Fail()
		t.Logf("TestGetTotalYearlyWeight failed. unexpected error: %s", err.Error())
	}

	if result[2022] != uint32(300) && result[2023] != uint32(1067) {
		t.Fail()
		t.Logf("TestGetTotalYearlyWeight failed. incorrect data returned")
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
