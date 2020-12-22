package dbutils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitConnection returns a client that connects to the db
func InitConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost"))
	// defer client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected to the db :success")
	}
	// todo - address potential leak with mongodb client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
