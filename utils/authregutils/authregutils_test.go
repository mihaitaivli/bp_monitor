package authregutils

import (
	"github.com/mihaitaivli/bp_monitor/graph/model"
	"testing"
)

func TestNewRegistrationInput (t *testing.T) {
	ri := NewRegistrationInput(TestUserInput)

	if ri.Email != TestUserInput.Email {
		t.Log("Incorrect value for NewRegistrationInput Email field")
		t.Fail()
	}

	if ri.Phone != nil {
		t.Log("Incorrect value for NewRegistrationInput Phone field")
		t.Fail()
	}
}

var TestUserInput = model.AddUserInput{
	Name:     "Some Name",
	Email:    "someone@test.com",
	Password: "testPassword",
	Phone:    nil,
	Dob:      nil,
}
