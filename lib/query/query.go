package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNewQueryBuilder() primitive.D {
	query := make(bson.D, 0)
	return query
}

func AND(cond []map[string]interface{}, checkExistance bool) primitive.D {
	query := make(bson.D, 0)
	for _, data := range cond {
		for k, v := range data {
			if checkExistance {
				query = append(query, EXISTS(k, true))
			}
			query = append(query, EQ(k, v))
		}
	}
	return query
}

func EQ(k string, v interface{}) primitive.E {
	return bson.E{Key: k, Value: bson.M{"$eq": v}}
}

func NEQ(k string, v interface{}) primitive.E {
	return bson.E{Key: k, Value: bson.M{"$ne": v}}
}

func EXISTS(k string, v interface{}) primitive.E {
	return bson.E{Key: k, Value: bson.M{"$exists": v}}
}

func OR([]map[string]interface{}) {
	/*
	   primitive.E{Key: "$or", Value: bson.A{bson.D{primitive.E{Key: "salesman_details.id", Value: salesmanId}}, bson.D{primitive.E{Key: "collection_status.salesman_details.id", Value: salesmanId}}}},
	*/
}
