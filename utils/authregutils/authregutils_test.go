package authregutils

import (
	"github.com/mihaitaivli/bp_monitor/graph/model"
	"testing"
)

var TEST_USER_INPUT = model.AddUserInput{
	Name:     "Some Name",
	Email:    "someone@test.com",
	Password: "testPassword",
	Phone:    nil,
	Dob:      nil,
}

func TestNewRegistrationInput (t *testing.T) {
	ri := NewRegistrationInput(TEST_USER_INPUT)

	if ri.Email != TEST_USER_INPUT.Email {
		t.Log("Incorrect value for NewRegistrationInput Email field")
		t.Fail()
	}

	if ri.Phone != nil {
		t.Log("Incorrect value for NewRegistrationInput Phone field")
		t.Fail()
	}
}