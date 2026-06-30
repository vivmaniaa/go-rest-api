package mDB

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Event struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bsjon:"_id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Location    string             `json:"location"`
	DateTime    time.Time          `json:"datetime"`
	UserId      int
}

func InsertEvent(e Event) error {
	collection := mongoClient.Database(db).Collection(collName)
	inserted, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		fmt.Println("Unable to insert to the database!")
		return err
	}

	fmt.Println("Inserted the record with ID: ", inserted.InsertedID)
	return nil
}

func InsertMany(events []Event) error {
	var newEvents = make([]any, len(events))
	for i, event := range events {
		newEvents[i] = event
	}
	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.InsertMany(context.TODO(), newEvents)
	if err != nil {
		fmt.Println("Unable to insert to the database!")
		return err
	}

	fmt.Println(result)
	return nil

}

func UpdateEvent(eventId string, e Event) error {
	id, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": e.Name, "description": e.Description, "location": e.Location, "datetime": e.DateTime}}

	collection := mongoClient.Database(db).Collection(collName)
	var updated *mongo.UpdateResult
	updated, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated the document with ID: ", updated.UpsertedID)
	return nil

}
func UpdateEvents(eventId string, e Event) error {
	id, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": e.Name, "description": e.Description, "location": e.Location, "datetime": e.DateTime}}

	collection := mongoClient.Database(db).Collection(collName)
	var result *mongo.UpdateResult
	result, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated record: ", result)
	return nil

}

func DeleteEvent(eventID string) error {
	id, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Println("Deleted record: ", result)
	return nil
}

func Find(eventName string) Event {
	var event Event
	filter := bson.D{{Key: "name", Value: eventName}}
	collection := mongoClient.Database(db).Collection(collName)
	err := collection.FindOne(context.TODO(), filter).Decode(&event)

	if err != nil {
		log.Fatal(err)
	}

	return event

}
func FindAll(eventName string) []Event {
	var events []Event
	filter := bson.D{{Key: "name", Value: eventName}}
	collection := mongoClient.Database(db).Collection(collName)
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(context.TODO(), &events)

	return events

}
func ListAll() []Event {
	var events []Event
	filter := bson.D{}
	collection := mongoClient.Database(db).Collection(collName)
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(context.TODO(), &events)

	return events

}

func DeleteAll() error {

	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	fmt.Printf("All the %d records are deleted", result.DeletedCount)
	return nil
}
