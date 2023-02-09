package models

import (
	"testing"
)

func TestWithExerciseId(t *testing.T) {
	expectedId := uint32(123)
	option := NewExerciseQueryOptions(WithExerciseId(expectedId))

	if option.Id == nil {
		t.Fail()
		t.Logf("TestWithExerciseId failed. expected: %v, got: nil", expectedId)
	}

	if *option.Id != expectedId {
		t.Fail()
		t.Logf("TestWithExerciseId failed. expected: %v, got: %v", expectedId, *option.Id)
	}

	if option.Title != nil {
		t.Fail()
		t.Logf("TestWithExerciseId failed: unexpected options present")
	}
}

func TestWithExerciseTitle(t *testing.T) {
	expectedTitle := "trainHeroic"
	option := NewExerciseQueryOptions(WithExerciseTitle(expectedTitle))

	if option.Title == nil {
		t.Fail()
		t.Logf("TestWithExerciseTitle failed. expected: %v, got: nil", expectedTitle)
	}

	if *option.Title != expectedTitle {
		t.Fail()
		t.Logf("TestWithExerciseTitle failed. expected: %v, got: %v", expectedTitle, *option.Title)
	}

	if option.Id != nil {
		t.Fail()
		t.Logf("TestWithExerciseTitle failed: unexpected options present")
	}
}
