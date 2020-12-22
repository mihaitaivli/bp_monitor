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

// InputIsValid returns a boolean reflecting the validation status of the input
func InputIsValid(input model.AddUserInput) error {

	// check if user exists
	userAlreadyExists, err := EmailAlreadyRegistered(input.Email)
	// check if email is valid
	emailIsValid := EmailIsValid(input.Email)

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
func EmailAlreadyRegistered(email string) (bool, error) {
	collection := client.Database("bp_log").Collection("users")
	filter := bson.D{{"email", email}}

	count, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// EmailIsValid superficially checks if the email entered is valid or not
func EmailIsValid(email string) bool {
	match, err := regexp.MatchString(`\w*\W*@\w+.\w+`, email)

	if err != nil {
		panic("invalid regex for email validation")
	}

	return match
}
