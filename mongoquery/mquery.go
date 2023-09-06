package mongoquery

import (
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	client     *mongo.Client
	database   string
	collection string
}

func New(mongoURI, database, collection string) (*MongoDBClient, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	return &MongoDBClient{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (c *MongoDBClient) Close() error {
	return c.client.Disconnect(context.Background())
}

func (c *MongoDBClient) Query(input string) ([]bson.M, error) {
	pairs := ParseInputString(input)

	query := bson.M{}
	for _, pair := range pairs {
		for key, value := range pair {
			query[key] = value
		}
	}

	coll := c.client.Database(c.database).Collection(c.collection)
	cursor, err := coll.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []bson.M
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Printf("Error decoding document: %v", err)
			continue
		}
		results = append(results, result)
	}

	return results, nil
}

func ParseInputString(input string) []map[string]interface{} {
	input = strings.ReplaceAll(input, "+", ",")
	parts := strings.Split(input, ",")
	var pairs []map[string]interface{}

	for _, part := range parts {
		subParts := strings.SplitN(part, ":", 2)
		if len(subParts) == 2 {
			key := strings.TrimSpace(subParts[0])
			value := strings.TrimSpace(subParts[1])
			if strings.Contains(value, "*") {
				value = strings.ReplaceAll(value, "*", ".*")
				value = "^" + value + "$"
				query := map[string]interface{}{key: primitive.Regex{Pattern: value}}
				pairs = append(pairs, query)
			} else {
				query := map[string]interface{}{key: value}
				pairs = append(pairs, query)
			}
		}
	}

	return pairs
}
