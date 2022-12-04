package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func registerProfile(uri, login, userID string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("userData")
	ProfileCollection := quickstartDatabase.Collection("Profile")

	ProfileCollection.InsertOne(ctx, bson.D{
		{Key: "FullName", Value: login},
		{Key: "Id", Value: userID},
		{Key: "Description", Value: "Default description"},
		{Key: "PhoneNB", Value: "none"},
		{Key: "Email", Value: "none"},
	})
	return true
}

func publishProfileUpdates(uri, endpoint, userID, data string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("userData")
	ProfileCollection := quickstartDatabase.Collection("Profile")

	filter := bson.D{{Key: "Id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: endpoint, Value: data}}}}
	_, err = ProfileCollection.UpdateOne(ctx, filter, update)

	return true
}

func getUserProfile(uri, userID string) primitive.M {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("userData")
	ProfileCollection := quickstartDatabase.Collection("Profile")

	var result bson.M
	querr := ProfileCollection.FindOne(context.TODO(), bson.D{{Key: "Id", Value: userID}}).Decode(&result)
	if querr != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if querr == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	// fmt.Printf("found document %v", result)

	return result
}

func searchUser(uri, username string) []primitive.M {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("userData")
	ProfileCollection := quickstartDatabase.Collection("Profile")

	filter := bson.D{{Key: "FullName", Value: username}}
	cursor, querr := ProfileCollection.Find(ctx, filter)

	if querr != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	if err != nil {
		log.Fatal(err)
	}

	return results
}
