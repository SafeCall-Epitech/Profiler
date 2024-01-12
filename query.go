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

type Contact struct {
	FullName    string   `bson:"FullName"`
	Id          string   `bson:"Id"`
	Description string   `bson:"Description"`
	PhoneNB     string   `bson:"PhoneNB"`
	Email       string   `bson:"Email"`
	Friends     []string `bson:"Friends"`
	Agenda      []Event  `bson:"Agenda"`
	ProfilePic  string   `bson:"ProfilePic"`
}

type Event struct {
	Guests    string `bson:"Guests"`
	Date      string `bson:"Date"`
	Subject   string `bson:"Subject"`
	Confirmed bool   `bson:"Confirmed"`
}

type Friends struct {
	Id       string  `bson:"Id"`
	Subject  string  `bson:"Subject"`
	Active   bool    `bson:"Active"`
	FullName *string `bson:"FullName"`
}

type Notification struct {
	Title   string `bson:"Title"`
	Content string `bson:"Content"`
	Status  bool   `bson:"Status"`
}

func registerProfile(uri, login, email string) bool {
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

	user := Contact{
		FullName:    login,
		Id:          login,
		Description: "Default description",
		PhoneNB:     "none",
		Email:       email,
		Friends:     []string{},
		Agenda:      []Event{},
		ProfilePic:  "default",
	}
	_, err = ProfileCollection.InsertOne(context.Background(), user)

	if err != nil {
		fmt.Println("Failed to insert contact:", err)
		return false
	}

	return true
}

func publishProfileUpdates(uri, endpoint, userID, data string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	quickstartDatabase := client.Database("userData")
	ProfileCollection := quickstartDatabase.Collection("Profile")

	filter := bson.D{{Key: "Id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: endpoint, Value: data}}}}
	result, err := ProfileCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return result.ModifiedCount == 1
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

	filter := bson.D{{Key: "FullName", Value: bson.D{
		{Key: "$regex", Value: primitive.Regex{Pattern: username, Options: "i"}},
	}}}
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

func addDelFriendAdd(uri, userID, action string, person Friends) bool {
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
	update := bson.D{{Key: action, Value: bson.D{{Key: "Friends", Value: person}}}}
	res, err := ProfileCollection.UpdateOne(ctx, filter, update)
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return err == nil
}

func GetFriends(uri, userID string) interface{} {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return "false"
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return "false"
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("userData")
	ProfileCollection := quickstartDatabase.Collection("Profile")
	filter := bson.D{{Key: "Id", Value: userID}}
	projection := bson.D{{Key: "Friends", Value: 1}}
	var result bson.M
	ProfileCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)

	return result["Friends"]
}

// works only if the id in the profile collection and the login in the login collection are the same !
func deleteUserProfile(uri, userID string) primitive.M {
	//init db + err handling
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
	LoginCollection := quickstartDatabase.Collection("loginInfo")

	var result bson.M
	//delete the profile colection
	querr := ProfileCollection.FindOneAndDelete(context.TODO(), bson.D{{Key: "Id", Value: userID}}).Decode(&result)
	if querr != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if querr == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}

	//delete the login collection
	querr2 := LoginCollection.FindOneAndDelete(context.TODO(), bson.D{{Key: "login", Value: userID}}).Decode(&result)
	if querr2 != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if querr2 == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	return result
}

func AddEvent(uri, dest string, event Event) bool {
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

	filter := bson.M{"Id": dest}
	update := bson.M{"$push": bson.M{"Agenda": event}}
	_, err = ProfileCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Failed to update contact:", err)
		return false
	}

	if err != nil {
		fmt.Println("Failed to insert contact:", err)
		return false
	}

	return true
}

func DelEvent(uri, dest, date string) bool {
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
	filter := bson.D{{Key: "Id", Value: dest}}
	update := bson.M{"$pull": bson.M{"Agenda": bson.M{"Date": date}}}
	_, err = ProfileCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Failed to update contact:", err)
		return false
	}

	return true
}

func GetNotificationsProfile(uri, userID string) interface{} {
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
	projection := bson.D{{Key: "Notifications", Value: 1}}
	var result bson.M
	ProfileCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)

	return result
}

func AddNotification(uri, UserID string, Notification Notification) bool {
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

	filter := bson.M{"Id": UserID}
	update := bson.M{"$push": bson.M{"Notifications": Notification}}
	_, err = ProfileCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Failed to update Notification:", err)
		return false
	}

	if err != nil {
		fmt.Println("Failed to insert notification:", err)
		return false
	}

	return true
}

func DelNotification(uri, UserID, Title string) bool {
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

	filter := bson.D{{Key: "Id", Value: UserID}}
	update := bson.M{"$pull": bson.M{"Notifications": bson.M{"Title": Title}}}
	found, err := ProfileCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Failed to delete notification:", err)
		return false
	}
	if found.MatchedCount == 0 {
		fmt.Println("Notification not found")
		return false
	} else {
		fmt.Println("Notification deleted notifications")
	}

	return true
}
