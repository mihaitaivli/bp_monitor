package authRegUtils

import (
	"context"
	"fmt"

	"github.com/mihaitaivli/bp_monitor/utils/dbUtils"
	"go.mongodb.org/mongo-driver/bson"
)

var client = dbUtils.InitConnection()

// RegistrationInput is a structure used for registration input validation
type RegistrationInput struct {
	Email       string  `json:"email"`
	RawPassword string  `json:"password"`
	Phone       *string `json:"phone,omitempty"`
}

// EmailAlreadyRegistered returns a boolean reflecting the existance in the db
// of an user with the same email address.
func (ri *RegistrationInput) EmailAlreadyRegistered(email string) bool {
	collection := client.Database("bp_log").Collection("users")
	filter := bson.D{{"email", ri.Email}}

	count, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		fmt.Println("Error while counting emails")
	}

	return count > 0
}
