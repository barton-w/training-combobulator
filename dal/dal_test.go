package dal

import (
	"testing"
	"training-combobulator/config"
	"training-combobulator/models"
)

func TestConnect(t *testing.T) {
	// Error conditions
	cfg, _ := config.NewConfig()
	cfg.Database.ExercisesDataPath = "/nodir/nofile.json"
	da, err := Connect(cfg.Database)

	if err == nil {
		t.Fail()
		t.Logf("TestConnect failed. error expected due to bad config")
	}

	if da != nil {
		t.Fail()
		t.Logf("TestConnect failed. expected nil, got %v", da)
	}

	cfg, _ = config.NewConfig()
	cfg.Database.UsersDataPath = "/nodir/nofile.json"
	da, err = Connect(cfg.Database)

	if err == nil {
		t.Fail()
		t.Logf("TestConnect failed. error expected due to bad config")
	}

	if da != nil {
		t.Fail()
		t.Logf("TestConnect failed. expected nil, got %v", da)
	}

	cfg, _ = config.NewConfig()
	cfg.Database.UsersDataPath = "/nodir/nofile.json"
	da, err = Connect(cfg.Database)

	if err == nil {
		t.Fail()
		t.Logf("TestConnect failed. error expected due to bad config")
	}

	if da != nil {
		t.Fail()
		t.Logf("TestConnect failed. expected nil, got %v", da)
	}

	cfg, _ = config.NewConfig()
	cfg.Database.WorkoutsDataPath = "/nodir/nofile.json"
	da, err = Connect(cfg.Database)

	if err == nil {
		t.Fail()
		t.Logf("TestConnect failed. error expected due to bad config")
	}

	if da != nil {
		t.Fail()
		t.Logf("TestConnect failed. expected nil, got %v", da)
	}

	// Success
	cfg, _ = config.NewConfig()
	cfg.Database = &config.Database{
		ExercisesDataPath: "../test/data/testExercises.json",
		UsersDataPath:     "../test/data/testUsers.json",
		WorkoutsDataPath:  "../test/data/testWorkouts.json",
	}
	da, err = Connect(cfg.Database)

	if err != nil {
		t.Fail()
		t.Logf("TestConnect failed. expected nil error, got: %s", err.Error())
	}

	if len(da.users) != 2 || len(da.exercises) != 2 || len(da.workouts) != 4 {
		t.Fail()
		t.Logf("TestConnect failed. datasources improperly parsed")
	}
}

func TestFindUsers(t *testing.T) {
	cfg, _ := config.NewConfig()
	cfg.Database = &config.Database{
		ExercisesDataPath: "../test/data/testExercises.json",
		UsersDataPath:     "../test/data/testUsers.json",
		WorkoutsDataPath:  "../test/data/testWorkouts.json",
	}
	da, _ := Connect(cfg.Database)

	expectedById := models.User{
		Id:        1,
		Firstname: "Alfred",
		Lastname:  "Music",
	}
	expectedByName := models.User{
		Id:        2,
		Firstname: "Train",
		Lastname:  "Heroic",
	}
	idResult := da.FindUsers(models.NewUserQueryOptions(models.WithUserId(1)))
	nameResult := da.FindUsers(models.NewUserQueryOptions(models.WithUserName("Train", "Heroic")))

	if len(idResult) != 1 || idResult[0] != expectedById {
		t.Fail()
		t.Logf("TestFindUsers failed. expected: %v, got: %v", expectedById, idResult[0])
	}

	if len(nameResult) != 1 || nameResult[0] != expectedByName {
		t.Fail()
		t.Logf("TestFindUsers failed. expected: %v, got: %v", expectedByName, nameResult[0])
	}
}

func TestFindExercises(t *testing.T) {
	cfg, _ := config.NewConfig()
	cfg.Database = &config.Database{
		ExercisesDataPath: "../test/data/testExercises.json",
		UsersDataPath:     "../test/data/testUsers.json",
		WorkoutsDataPath:  "../test/data/testWorkouts.json",
	}
	da, _ := Connect(cfg.Database)

	expectedById := models.Exercise{
		Id:    3,
		Title: "Bicep Curls",
	}
	expectedByTitle := models.Exercise{
		Id:    4,
		Title: "Tractor Pulls",
	}
	idResult := da.FindExercises(models.NewExerciseQueryOptions(models.WithExerciseId(3)))
	titleResult := da.FindExercises(models.NewExerciseQueryOptions(models.WithExerciseTitle("Tractor Pulls")))

	if len(idResult) != 1 || idResult[0] != expectedById {
		t.Fail()
		t.Logf("TestFindExercises failed. expected: %v, got: %v", expectedById, idResult[0])
	}

	if len(titleResult) != 1 || titleResult[0] != expectedByTitle {
		t.Fail()
		t.Logf("TestFindExercises failed. expected: %v, got: %v", expectedByTitle, titleResult[0])
	}
}

func TestFindWorkouts(t *testing.T) {
	cfg, _ := config.NewConfig()
	cfg.Database = &config.Database{
		ExercisesDataPath: "../test/data/testExercises.json",
		UsersDataPath:     "../test/data/testUsers.json",
		WorkoutsDataPath:  "../test/data/testWorkouts.json",
	}
	da, _ := Connect(cfg.Database)

	userIdResult := da.FindWorkouts(models.NewWorkoutQueryOptions(models.WithWorkoutUserId(1)))
	if len(userIdResult) != 2 {
		t.Fail()
		t.Logf("TestFindWorkouts failed. expected: 2 workouts, got: %v", len(userIdResult))
	}

	for _, workout := range userIdResult {
		if len(workout.Blocks) != 2 {
			t.Fail()
			t.Logf("TestFindWorkouts failed. expected: 2 blocks, got: %v", len(workout.Blocks))
		}
		for _, block := range workout.Blocks {
			if len(block.Sets) != 2 {
				t.Fail()
				t.Logf("TestFindWorkouts failed. expected: 2 set, got: %v", len(block.Sets))
			}
		}
	}

	exIdResult := da.FindWorkouts(models.NewWorkoutQueryOptions(models.WithWorkoutExerciseId(4)))
	if len(exIdResult) != 4 {
		t.Fail()
		t.Logf("TestFindWorkouts failed. expected: 4 workouts, got: %v", len(exIdResult))
	}

	for _, workout := range exIdResult {
		if len(workout.Blocks) != 1 {
			t.Fail()
			t.Logf("TestFindWorkouts failed. expected: 1 block, got: %v", len(workout.Blocks))
		}
		for _, block := range workout.Blocks {
			if len(block.Sets) != 2 {
				t.Fail()
				t.Logf("TestFindWorkouts failed. expected: 2 set, got: %v", len(block.Sets))
			}
		}
	}

	comboResult := da.FindWorkouts(models.NewWorkoutQueryOptions(models.WithWorkoutUserId(2), models.WithWorkoutExerciseId(3)))
	if len(comboResult) != 2 {
		t.Fail()
		t.Logf("TestFindWorkouts failed. expected: 4 workouts, got: %v", len(comboResult))
	}

	for _, workout := range comboResult {
		if len(workout.Blocks) != 1 {
			t.Fail()
			t.Logf("TestFindWorkouts failed. expected: 1 block, got: %v", len(workout.Blocks))
		}
		for _, block := range workout.Blocks {
			if len(block.Sets) != 2 {
				t.Fail()
				t.Logf("TestFindWorkouts failed. expected: 2 set, got: %v", len(block.Sets))
			}
		}
	}
}

func TestFilterBlocksByExerciseId(t *testing.T) {
	input := []models.Block{
		{
			ExerciseId: 8675,
			Sets:       []models.Set{},
		},
		{
			ExerciseId: 42,
			Sets:       []models.Set{},
		},
	}

	result := filterBlocksByExerciseId(input, uint32(42))
	if len(result) != 1 || result[0].ExerciseId != uint32(42) {
		t.Fail()
		t.Logf("TestFilterBlocksByExerciseId failed. unexpected return value")
	}
}
