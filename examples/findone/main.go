package main

import (
	"fmt"
	"libraries/go-mongo-query/lib/mongodb"
	"libraries/go-mongo-query/lib/queryer"
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

	mongoQueryer := queryer.New()
	q := mongoQueryer.SetConnection(connection).GetQueryer("test", "test_collection")
	defer mongoQueryer.Close()

	// qu := query.GetNewQueryBuilder()
	// qu = append(qu, query.EXISTS("first_name", true))
	// qu = append(qu, query.EQ("first_name", "Catarina"))
	whereClause := make(map[string]interface{})
	whereClause["first_name"] = "Catarina"
	err := q.FindOneMatching(whereClause, data)
	fmt.Println(err)
	fmt.Println(data.FirstName, data.LastName)

	notwhereClause := make(map[string]interface{})
	notwhereClause["gender"] = "Polygender"
	err = q.FindOneNotMatching(notwhereClause, data)
	fmt.Println(err)
	fmt.Println(data.FirstName, data.LastName)

	/*****************************************************/
	//Findone with contional statements
	/****************************************************/

	// whereClass1 := make(map[string]interface{})
	// whereClass["first_name"] = "Catarina"

	// whereClass2 := make(map[string]interface{})
	// whereClass["gender"] = "Polygender"

	// q.FindOne(whereClass, data)
	// //res.Decode(data)
	// fmt.Println(data.FirstName, data.LastName)

}
