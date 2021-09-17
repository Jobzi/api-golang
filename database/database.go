package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	database="todo-app"
)

//GetCollection send name collection recive collection
func GetCollection(collection string) (*mongo.Collection){
	/*
		Connect to my cluster
    */
	connectionString:="mongodb+srv://dbUser:Ib3IamQHxWxt4ZOq@cluster0.mjeri.mongodb.net/todo-app-test?retryWrites=true&w=majority"
    client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
    if err != nil {
        log.Fatal(err)
    }
    ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//defer client.Disconnect(ctx)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
	return client.Database(database).Collection(collection)
}