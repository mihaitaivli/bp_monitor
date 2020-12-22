package authregutils

import (
	"testing"
)

func TestEmailIsValid (t *testing.T) {
	faultyEmails := []string{
		"",
		"a",
		"@",
		"a@",
		"a@.",
		"a@.com",
		"a@a",
		"a@a.",
	}
	validEmail := "a@a.com"

	for _, email := range faultyEmails {
		emailValid := emailIsValid(email)
		if emailValid == true  {
			t.Logf("Faulty email %s should be marked as invalid", email)
			t.Fail()
		}
	}

	emailValid := emailIsValid(validEmail)
	if emailValid != true  {
		t.Logf("valid email %s should pass validation test", validEmail)
		t.Fail()
	}
}

//var TestUserInput = model.AddUserInput{
//	Name:     "Some Name",
//	Email:    "someone@test.com",
//	Password: "testPassword",
//	Phone:    nil,
//	Dob:      nil,
//}
