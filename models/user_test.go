package models

import (
	"testing"
)

func TestWithUserId(t *testing.T) {
	expectedId := uint32(123)
	option := NewUserQueryOptions(WithUserId(expectedId))

	if option.Id == nil {
		t.Fail()
		t.Logf("TestWithUserId failed. expected: %v, got: nil", expectedId)
	}

	if *option.Id != expectedId {
		t.Fail()
		t.Logf("TestWithUserId failed. expected: %v, got: %v", expectedId, *option.Id)
	}

	if option.Firstname != nil || option.Lastname != nil {
		t.Fail()
		t.Logf("TestWithUserId failed: unexpected options present")
	}
}

func TestWithUserName(t *testing.T) {
	expectedFirst := "training"
	expectedLast := "peaks"
	option := NewUserQueryOptions(WithUserName(expectedFirst, expectedLast))

	if option.Firstname == nil {
		t.Fail()
		t.Logf("TestWithUserName failed. expected: %v, got: nil", expectedFirst)
	}

	if option.Lastname == nil {
		t.Fail()
		t.Logf("TestWithUserName failed. expected: %v, got: nil", expectedLast)
	}

	if *option.Firstname != expectedFirst {
		t.Fail()
		t.Logf("TestWithUserName failed. expected: %v, got: %v", expectedFirst, *option.Firstname)
	}

	if *option.Lastname != expectedLast {
		t.Fail()
		t.Logf("TestWithUserName failed. expected: %v, got: %v", expectedLast, *option.Lastname)
	}

	if option.Id != nil {
		t.Fail()
		t.Logf("TestWithUserName failed: unexpected options present")
	}
}
