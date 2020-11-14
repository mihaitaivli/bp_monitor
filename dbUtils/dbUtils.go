package dbUtils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// InitConnection returns a client that connects to the db
func InitConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost"))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected to the db :success")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
