package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mihaitaivli/bp_monitor/graph/generated"
	"github.com/mihaitaivli/bp_monitor/graph/model"
	"github.com/mihaitaivli/bp_monitor/utils/authregutils"
	"github.com/mihaitaivli/bp_monitor/utils/dbutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) AddUser(ctx context.Context, input model.AddUserInput) (*string, error) {
	collection := client.Database("bp_log").Collection("users")

	// registration checks
	registrationInput := authregutils.NewRegistrationInput(input)
	inputIsValid, error := registrationInput.InputIsValid()


	insertUserResult, err := collection.InsertOne(context.Background(), input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	insertedID := insertUserResult.InsertedID.(primitive.ObjectID).Hex()

	return &insertedID, nil
}

func (r *mutationResolver) AddRecord(ctx context.Context, input model.NewRecord) (*string, error) {
	collection := client.Database("bp_log").Collection("records")

	insertRecordResult, err := collection.InsertOne(context.Background(), input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	insertedID := insertRecordResult.InsertedID.(primitive.ObjectID).Hex()

	return &insertedID, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	collection := client.Database("bp_log").Collection("users")

	objID, error := primitive.ObjectIDFromHex(id)
	if error != nil {
		log.Println(error)
		return nil, error
	}

	filter := bson.M{"_id": objID}
	var result MongoFields

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user := model.User{
		ID:   result.ID.Hex(),
		Name: result.Name,
	}
	return &user, nil
}

func (r *queryResolver) Record(ctx context.Context, id string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Records(ctx context.Context, where *model.RecordsWhere, sortBy *model.RecordsSortBy, paginate *model.Pagination) ([]*model.Record, error) {
	record := model.Record{
		ID:        "123",
		Systolic:  120,
		Diastolic: 70,
	}

	records := []*model.Record{&record}
	return records, nil
}

func (r *userResolver) Records(ctx context.Context, obj *model.User) ([]*model.Record, error) {
	fmt.Println(ctx)

	record := model.Record{
		ID:        "123456789",
		Systolic:  140,
		Diastolic: 75,
	}

	records := []*model.Record{&record}
	return records, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var client = dbutils.InitConnection()

type MongoFields struct {
	ID   primitive.ObjectID `bson:"_id, omitempty"`
	Name string             `bson:"name"`
}
