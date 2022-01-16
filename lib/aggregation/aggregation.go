package aggregation

import "go.mongodb.org/mongo-driver/mongo"

/*
queryer := o.dbConn.Collection("order")
	query := make(bson.D, 1)
	query = append(query, bson.E{Key: "bill_no", Value: bson.M{"$exists": true}})
	query = append(query, bson.E{Key: "bill_no", Value: bson.M{"$eq": billNo}})

	o.l.Debugf("Query : %v", query)
	result := dtos.Order{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	o.l.Debug("Before Q")
	err := queryer.FindOne(ctx, query).Decode(&result)

	if err != nil {
		o.l.Errorf("Err : %v", err.Error())
		return nil, err
	}
	o.l.Debug("After Q : %v", result)


*/
type mongoQueryer struct {
	dbconn     mongo.Client
	dataBase   mongo.Database
	collection mongo.Collection
}

// func Where(field string, value string) []bson.D {

// }
