package main

import (
	"fmt"
	"libraries/go-mongo-query/lib/mongodb"
	"libraries/go-mongo-query/lib/query"
)

type Sample struct {
	ID        string `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Gender    string `bson:"gender"`
	IPAddress string `bson:"ip_address"`
	CreatedAt string `bson:"created_at"`
}

func main() {
	fmt.Println("Hello World")
	mongodb.Connect("mongodb://localhost:27017")
	connection := mongodb.NewConnection()
	// // Declare Context type object for managing multiple API requests
	// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// // Access a MongoDB collection through a database
	// col := connection.Database("test").Collection("test_collection")
	data := &Sample{}
	whereClass := make(map[string]interface{})
	whereClass["first_name"] = "Catarina"
	mongoQueryer := query.New()
	q := mongoQueryer.Connection(connection).GetQueryer("test", "test_collection")

	q.FindOne(whereClass, data)
	//res.Decode(data)
	fmt.Println(data.FirstName, data.LastName)

}
