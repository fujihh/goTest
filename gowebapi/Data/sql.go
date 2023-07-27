package sql

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func main() {
// 	var(
// 		client 		*mongo.Client
// 		err       	 error
// 		db         	*mongo.Database
// 		collection 	*mongo.Collection
// 	)

// 	 //1.建立连接
// 	 if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
// 		fmt.Print(err)
// 		return
// 	 }
// 	 //2.选择数据库 my_db
// 	 db = client.Database("my_db")

//		 //3.选择表 my_collection
//		 collection = db.Collection("my_collection")
//		 collection = collection
//	}
type Person struct {
	Name  string
	Age   int
	Email string
}
type Trainer struct {
	Name string
	Age  int
	City string
}

func sqlconnect(ctx context.Context) (*mongo.Client, error) {
	// Set client options
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	// insert
	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// find
	var result Trainer
	err = collection.FindOne(context.TODO(), Trainer{"Ash", 10, "Pallet Town"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found a single document: ", result)
	return client, err
}
