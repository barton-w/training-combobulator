package models

import (
	"testing"
	"time"
)

func TestWithWorkoutUserId(t *testing.T) {
	expectedId := uint32(123)
	option := NewWorkoutQueryOptions(WithWorkoutUserId(expectedId))

	if option.UserId == nil {
		t.Fail()
		t.Logf("TestWithWorkoutUserId failed. expected: %v, got: nil", expectedId)
	}

	if *option.UserId != expectedId {
		t.Fail()
		t.Logf("TestWithWorkoutUserId failed. expected: %v, got: %v", expectedId, *option.UserId)
	}

	if option.ExerciseId != nil {
		t.Fail()
		t.Logf("TestWithWorkoutUserId failed: unexpected options present")
	}
}

func TestWithWorkoutExerciseId(t *testing.T) {
	expectedId := uint32(8675)
	option := NewWorkoutQueryOptions(WithWorkoutExerciseId(expectedId))

	if option.ExerciseId == nil {
		t.Fail()
		t.Logf("TestWithWorkoutExerciseId failed. expected: %v, got: nil", expectedId)
	}

	if *option.ExerciseId != expectedId {
		t.Fail()
		t.Logf("TestWithWorkoutExerciseId failed. expected: %v, got: %v", expectedId, *option.ExerciseId)
	}

	if option.UserId != nil {
		t.Fail()
		t.Logf("TestWithWorkoutExerciseId failed: unexpected options present")
	}
}

func TestMultiOption(t *testing.T) {
	expectedUserId := uint32(666)
	expectedExerciseId := uint32(999)
	option := NewWorkoutQueryOptions(WithWorkoutUserId(expectedUserId), WithWorkoutExerciseId(expectedExerciseId))

	if option.UserId == nil {
		t.Fail()
		t.Logf("TestMultiOption failed. expected: %v, got: nil", expectedUserId)
	}

	if option.ExerciseId == nil {
		t.Fail()
		t.Logf("TestMultiOption failed. expected: %v, got: nil", expectedExerciseId)
	}

	if *option.ExerciseId != expectedExerciseId {
		t.Fail()
		t.Logf("TestMultiOption failed. expected: %v, got: %v", expectedExerciseId, *option.ExerciseId)
	}

	if *option.UserId != expectedUserId {
		t.Fail()
		t.Logf("TestMultiOption failed. expected: %v, got: %v", expectedUserId, *option.UserId)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	wd := &WorkoutDateTime{}
	expected := time.Date(2016, 1, 25, 17, 7, 33, 0, time.UTC)
	valid := []byte(`"2016-01-25 17:07:33"`)
	invalid := []byte(`"23-15-00385 34:12"`)
	err := wd.UnmarshalJSON(valid)

	if err != nil {
		t.Fail()
		t.Logf("TestUnmarshalJSON failed with error: %s", err.Error())
	}

	if wd.Time != expected {
		t.Fail()
		t.Logf("TestUnmarshalJSON failed. expected: %v, got: %v", expected, wd.Time)
	}

	err = wd.UnmarshalJSON(invalid)

	if err == nil {
		t.Fail()
		t.Logf("TestUnmarshalJSON failed, error expected")
	}

}
