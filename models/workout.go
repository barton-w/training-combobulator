package models

type Set struct {
	Reps   uint32  `json:"reps"`
	Weight *uint32 `json:"weight"`
}

type Block struct {
	ExerciseId uint32 `json:"exercise_id"`
	Sets       []Set
}

type Workout struct {
	UserId            uint32 `json:"user_id"`
	DatetimeCompleted string `json:"datetime_completed"`
	Blocks            []Block
}
