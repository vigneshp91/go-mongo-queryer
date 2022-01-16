package queryer

import (
	"context"
	"libraries/go-mongo-query/lib/mongodb"
	"libraries/go-mongo-query/lib/query"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoQueryerBuilder Interface for all mongo query functions
type MongoQueryerBuilder interface {
	SetConnection(conn *mongodb.Connection) MongoQueryerBuilder
	WithContext(context.Context) MongoQueryerBuilder //optional
	Close()
	GetQueryer(db string, collection string) MongoQueryerBuilder
	FindOne(req primitive.D, result interface{}, opts ...*options.FindOneOptions) error
	FindOneMatching(req map[string]interface{}, result interface{}, opts ...*options.FindOneOptions) error
	FindOneNotMatching(req map[string]interface{}, result interface{}, opts ...*options.FindOneOptions) error
}

type queryerBuilder struct {
	dbConn  *mongodb.Connection
	ctx     context.Context
	cancel  context.CancelFunc
	queryer *mongo.Collection
}

func (cb *queryerBuilder) SetConnection(conn *mongodb.Connection) MongoQueryerBuilder {
	cb.dbConn = conn
	return cb
}

func (cb *queryerBuilder) WithContext(ctx context.Context) MongoQueryerBuilder {
	cb.ctx = ctx
	return cb
}

func (cb *queryerBuilder) GetQueryer(db string, collection string) MongoQueryerBuilder {
	cb.queryer = cb.dbConn.Database(db).Collection(collection)
	return cb
}

func (cb *queryerBuilder) Close() {
	cb.cancel()
}

func (cb *queryerBuilder) FindOne(cond primitive.D, result interface{}, opts ...*options.FindOneOptions) error {
	// q := make(bson.D, 0)
	//	q := query.AND([]map[string]interface{}{cond}, true)
	// for k, v := range cond {
	// 	query = append(query, bson.E{Key: k, Value: bson.M{"$exists": true}})
	// 	query = append(query, bson.E{Key: k, Value: bson.M{"$eq": v}})
	// }
	return cb.queryer.FindOne(cb.ctx, cond, opts...).Decode(result)
}

func (cb *queryerBuilder) FindOneMatching(cond map[string]interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	q := make(bson.D, 0)
	for k, v := range cond {
		q = append(q, query.EXISTS(k, true))
		q = append(q, query.EQ(k, v))
	}
	return cb.queryer.FindOne(cb.ctx, q, opts...).Decode(result)
}

func (cb *queryerBuilder) FindOneNotMatching(cond map[string]interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	q := make(bson.D, 0)
	for k, v := range cond {
		// q = append(q, query.EXISTS(k, true))
		q = append(q, query.NEQ(k, v))
	}
	return cb.queryer.FindOne(cb.ctx, q, opts...).Decode(result)
}

//New Returns new instance of Mongo ConnectionQueryer
func New() MongoQueryerBuilder {
	q := &queryerBuilder{}
	q.ctx, q.cancel = context.WithTimeout(context.Background(), 15*time.Second)
	return q
}
