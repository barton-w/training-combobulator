# Welcome to the Training Combobulator!

## Packages overview

#### main
Entrypoint to the app, which initializes config, connects to the database,
and invokes handlers to answer user, exercise, and workout queries

#### config
Defines and provides a configuration object for the app, reading configuration kvs from the environment

#### dal (data access layer)
Simulates a database connection by reading-in the various JSON files.
The data access layer implements the interface required by handlers for accessing data.

The intent was to create a simple abstraction layer which would allow the data access layer
to change to another datastore while still implementing the simple interface.

#### models
Defines data structures used in the app, and functional query-options for 
flexiblity when querying the data access layer

#### handlers
Handlers contain business logic and calculations to solve the problems in the assignment, and beyond.
The intent is that handlers could be converted to http route handlers, and parse request data
to fulfil a variety of queries.

While these handlers do answer the questions in the assignment, they are user and exercise agnostic, 
and can also provide year and monthly workout aggregations for UIs, consumers, etc.

## How to build and run

#### Using Docker
1. `docker build -t training-combobulator .`
2. `docker run --rm --env-file .env training-combobulator`

#### If you don't have Docker, but do have Go 1.20 installed
Earlier versions of Go should also work, however I've only tested using Go 1.20.

`export EXERCISES_PATH=dal/data/exercises.json USERS_PATH=dal/data/users.json WORKOUTS_PATH=dal/data/workouts.json && go run cmd/main.go`

## How to run tests
`go test -v ./...`

Here's the output from my latest test-run:

```
?   	training-combobulator/cmd	[no test files]
=== RUN   TestNewConfig
--- PASS: TestNewConfig (0.00s)
PASS
ok  	training-combobulator/config	0.160s
=== RUN   TestConnect
--- PASS: TestConnect (0.00s)
=== RUN   TestFindUsers
--- PASS: TestFindUsers (0.00s)
=== RUN   TestFindExercises
--- PASS: TestFindExercises (0.00s)
=== RUN   TestFindWorkouts
--- PASS: TestFindWorkouts (0.00s)
=== RUN   TestFilterBlocksByExerciseId
--- PASS: TestFilterBlocksByExerciseId (0.00s)
PASS
ok  	training-combobulator/dal	0.334s
=== RUN   TestFindOneUser
--- PASS: TestFindOneUser (0.00s)
=== RUN   TestFindOneExercise
--- PASS: TestFindOneExercise (0.00s)
=== RUN   TestGetExerciseTotalWeight
--- PASS: TestGetExerciseTotalWeight (0.00s)
=== RUN   TestGetUserYearlyExerciseTotals
--- PASS: TestGetUserYearlyExerciseTotals (0.00s)
=== RUN   TestGetMaxWeightMonths
--- PASS: TestGetMaxWeightMonths (0.00s)
=== RUN   TestGetTotalYearlyWeight
--- PASS: TestGetTotalYearlyWeight (0.00s)
=== RUN   TestGetUserPR
--- PASS: TestGetUserPR (0.00s)
=== RUN   TestMaxWeight
--- PASS: TestMaxWeight (0.00s)
=== RUN   TestBlockTotals
--- PASS: TestBlockTotals (0.00s)
PASS
ok  	training-combobulator/handlers	0.249s
=== RUN   TestWithExerciseId
--- PASS: TestWithExerciseId (0.00s)
=== RUN   TestWithExerciseTitle
--- PASS: TestWithExerciseTitle (0.00s)
=== RUN   TestWithUserId
--- PASS: TestWithUserId (0.00s)
=== RUN   TestWithUserName
--- PASS: TestWithUserName (0.00s)
=== RUN   TestWithWorkoutUserId
--- PASS: TestWithWorkoutUserId (0.00s)
=== RUN   TestWithWorkoutExerciseId
--- PASS: TestWithWorkoutExerciseId (0.00s)
=== RUN   TestMultiOption
--- PASS: TestMultiOption (0.00s)
=== RUN   TestUnmarshalJSON
--- PASS: TestUnmarshalJSON (0.00s)
PASS
ok  	training-combobulator/models	0.414s
```