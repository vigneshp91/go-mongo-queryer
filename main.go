package main

import (
	"context"
	"fmt"
	"libraries/go-mongo-query/lib/mongodb"
	"libraries/go-mongo-query/lib/query"
	"time"
)

func main() {
	fmt.Println("Hello World")
	mongodb.Connect("mongodb://localhost:27017")
	connection := mongodb.NewConnection()
	// // Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// // Access a MongoDB collection through a database
	// col := connection.Database("test").Collection("test_collection")
	whereClass := make(map[string]interface{})
	whereClass["first_name"] = "Vanny"
	mongoQueryer := query.New()
	res, err := mongoQueryer.Connection(connection).
		Database("test").
		Collection("test_collection").
		WithContext(ctx).Query().FindOne(whereClass)

}
