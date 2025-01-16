package connection

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collect *mongo.Collection

func init() {
	const Url = "mongodb://localhost:27017"
	const dbname = "Netflix"
	const colname = "Watchlist"

	clientoption := options.Client().ApplyURI(Url)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		fmt.Printf("Error while connecting %v", err)
		return
	}

	fmt.Println("Connection successful")

	Collect = client.Database(dbname).Collection(colname)
	fmt.Println("collection to database established")
}
