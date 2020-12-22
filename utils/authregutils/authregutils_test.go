package authregutils

import (
	"testing"
)

func TestEmailIsValid (t *testing.T) {
	for _, email := range faultyEmailsList {
		emailValid := EmailIsValid(email)
		if emailValid == true  {
			t.Logf("Faulty email %s should be marked as invalid", email)
			t.Fail()
		}
	}

	emailValid := EmailIsValid(validEmailExample)
	if emailValid != true  {
		t.Logf("valid email %s should pass validation test", validEmailExample)
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
