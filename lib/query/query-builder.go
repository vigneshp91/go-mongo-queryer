package query

import (
	"context"
	"libraries/go-mongo-query/lib/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

//MongoQueryBuilder Interface for all mongo query functions
type MongoQueryBuilder interface {
	Connection(conn *mongodb.Connection) MongoQueryBuilder
	Collection(string) MongoQueryBuilder
	Database(string) MongoQueryBuilder
	WithContext(context.Context) MongoQueryBuilder
	Query() MongoQueryBuilder
	FindOne(map[string]interface{}) (*mongo.SingleResult, error)
}

type queryBuilder struct {
	dbConn     *mongodb.Connection
	database   string
	collection string
	ctx        context.Context
	queryer    *mongo.Collection
}

func (cb *queryBuilder) Connection(conn *mongodb.Connection) MongoQueryBuilder {
	cb.dbConn = conn
	return cb
}

func (cb *queryBuilder) Database(database string) MongoQueryBuilder {
	cb.database = database
	return cb
}

func (cb *queryBuilder) Collection(collection string) MongoQueryBuilder {
	cb.collection = collection
	return cb
}

func (cb *queryBuilder) WithContext(ctx context.Context) MongoQueryBuilder {
	cb.ctx = ctx
	return cb
}

func (cb *queryBuilder) Query() MongoQueryBuilder {
	// mongodb.Connect("mongodb://localhost:27017")
	// connection := mongodb.NewConnection()
	cb.queryer = cb.dbConn.Database(cb.database).Collection(cb.collection)
	return cb
}

func (cb *queryBuilder) FindOne(map[string]interface{}) (*mongo.SingleResult, error) {

}

//New Returns new instance of Mongo Queryer
func New() MongoQueryBuilder {
	return &queryBuilder{}
}
