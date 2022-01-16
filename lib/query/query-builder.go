package query

import (
	"context"
	"libraries/go-mongo-query/lib/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoQueryBuilder Interface for all mongo query functions
type MongoQueryBuilder interface {
	Connection(conn *mongodb.Connection) MongoQueryBuilder
	// Collection(string) MongoQueryBuilder
	// Database(string) MongoQueryBuilder
	WithContext(context.Context) MongoQueryBuilder
	GetQueryer(db string, collection string) MongoQueryBuilder
	FindOne(req map[string]interface{}, result interface{}) error
}

type queryBuilder struct {
	dbConn *mongodb.Connection
	// database   string
	// collection string
	ctx     context.Context
	cancel  context.CancelFunc
	queryer *mongo.Collection
}

func (cb *queryBuilder) Connection(conn *mongodb.Connection) MongoQueryBuilder {
	cb.dbConn = conn
	return cb
}

// func (cb *queryBuilder) Database(database string) MongoQueryBuilder {
// 	cb.database = database
// 	return cb
// }

// func (cb *queryBuilder) Collection(collection string) MongoQueryBuilder {
// 	cb.collection = collection
// 	return cb
// }

func (cb *queryBuilder) WithContext(ctx context.Context) MongoQueryBuilder {
	cb.ctx = ctx
	return cb
}

func (cb *queryBuilder) GetQueryer(db string, collection string) MongoQueryBuilder {
	// mongodb.Connect("mongodb://localhost:27017")
	// connection := mongodb.NewConnection()
	cb.queryer = cb.dbConn.Database(db).Collection(collection)
	return cb
}

func (cb *queryBuilder) FindOne(cond map[string]interface{}, result interface{}) error {
	query := make(bson.D, 0)

	for k, v := range cond {
		query = append(query, bson.E{Key: k, Value: bson.M{"$exists": true}})
		query = append(query, bson.E{Key: k, Value: bson.M{"$eq": v}})
	}
	defer cb.cancel()
	return cb.queryer.FindOne(cb.ctx, query).Decode(result)
}

//New Returns new instance of Mongo ConnectionQueryer
func New() MongoQueryBuilder {
	q := &queryBuilder{}
	q.ctx, q.cancel = context.WithTimeout(context.Background(), 15*time.Second)
	return q
}
