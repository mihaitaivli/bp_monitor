package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mihaitaivli/bp_monitor/dbUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mihaitaivli/bp_monitor/graph/generated"
	"github.com/mihaitaivli/bp_monitor/graph/model"
)

var client = dbUtils.InitConnection()

func (r *mutationResolver) AddUser(ctx context.Context, input model.NewUser) (*string, error) {
	defer client.Disconnect(context.Background())
	collection := client.Database("bp_log").Collection("users")

	insertUserResult, err := collection.InsertOne(context.Background(), input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	insertedID := insertUserResult.InsertedID.(primitive.ObjectID).String()
	log.Println("Successfully addes user", insertedID)
	//return &insertedId, nil
	return &insertedID, nil
}

func (r *mutationResolver) AddRecord(ctx context.Context, input model.NewRecord) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Record(ctx context.Context, id string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Records(ctx context.Context, where model.RecordsWhere, sortBy *model.RecordsSortBy, paginate *model.Pagination) ([]*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
