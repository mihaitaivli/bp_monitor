package authregutils

import (
	"context"
	"fmt"
	"github.com/mihaitaivli/bp_monitor/graph/model"
	"regexp"

	"github.com/mihaitaivli/bp_monitor/utils/dbutils"
	"go.mongodb.org/mongo-driver/bson"
)

var client = dbutils.InitConnection()

// RegistrationInput is a structure used for registration input validation
type RegistrationInput struct {
	Email       string
	RawPassword string
	Phone       *string
}

// NewRegistrationInput initializes a new RegistrationInput struct and returns it.
func NewRegistrationInput(input model.AddUserInput) *RegistrationInput {
	return &RegistrationInput{
		Email: input.Email,
		RawPassword: input.Password,
		Phone: input.Phone,
	}
}

// InputIsValid returns a boolean reflecting the validation status of the input
func (ri *RegistrationInput) InputIsValid() error {

	// check if user exists
	userAlreadyExists, err := ri.EmailAlreadyRegistered()
	// check if email is valid
	emailIsValid := ri.EmailIsValid()

	if err != nil {
		return err
	}

	if !emailIsValid {
		return fmt.Errorf("email is invalid")
	}

	if userAlreadyExists {
		return fmt.Errorf("email already registered")
	}

	return nil
}

// EmailAlreadyRegistered returns a boolean reflecting the existence in the db
// of an user with the same email address.
func (ri *RegistrationInput) EmailAlreadyRegistered() (bool, error) {
	collection := client.Database("bp_log").Collection("users")
	filter := bson.D{{"email", ri.Email}}

	count, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// EmailIsValid superficially checks if the email entered is valid or not
func (ri *RegistrationInput) EmailIsValid() bool {
	match, err := regexp.MatchString(`\w*\W*@\w*.\w`, ri.Email)

	if err != nil {
		panic("invalid regex for email validation")
	}

	return match
}
