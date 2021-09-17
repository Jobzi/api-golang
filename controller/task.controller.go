package task_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	d "API/database"
	m "API/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

const (
	collectionName="tasks"
)

func Create(task m.TaskSchema) error {
	//fmt.Print(tasks)
	return nil
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	collection := d.GetCollection(collectionName)
	var tasks m.AllTasks

	filter := bson.D{}
	cur,err:=collection.Find(ctx,filter)
	if err != nil { log.Fatal(err) }
	//defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result m.TaskSchema
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		tasks = append(tasks, &result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PostTask(w http.ResponseWriter, r *http.Request){
	var newTask m.TaskSchema
	
	collection := d.GetCollection(collectionName)
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "INSER VALID TASK")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.Date = time.Now()
	_, err = collection.InsertOne(ctx, newTask)
	if err != nil {
		log.Fatal(err)
	}
	
	//sento to 
	w.Header().Set("Content-Type", "application/json") //cabecera
	w.WriteHeader(http.StatusCreated)                  //Si el tipo de dato fue correcto
	json.NewEncoder(w).Encode(newTask)                 //contendido

}