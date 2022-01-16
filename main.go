package main

import (
	"context"
	"fmt"
	"libraries/go-mongo-query/lib/mongodb"
)

func main() {
	fmt.Println("Hello World")
	mongodb.Connect("mongodb://localhost:27017")
	connection := mongodb.NewConnection()
	defer connection.Disconnect(context.Background())

}
