package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
		{Key: "email", Value: "none"},
	})
	return true
}

func publishDescription(uri, endpoint, userID, data string) bool {
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

// func GetUsers(uri string) []bson.M {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}
// 	defer client.Disconnect(ctx)

// 	quickstartDatabase := client.Database("userData")
// 	ProfileCollection := quickstartDatabase.Collection("loginInfo")

// 	cursor, err := ProfileCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var users []bson.M
// 	if err = cursor.All(ctx, &users); err != nil {
// 		log.Fatal(err)
// 	}

// 	return users
// }
