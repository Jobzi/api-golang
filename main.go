// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// // Types
// type taskSchema struct {
// 	ID      int    `json:"ID"`
// 	Name    string `json:"Name"`
// 	Content string `json:"Content"`
// }

// type allTasks []taskSchema

// // Persistence
// var tasks = allTasks{
// 	{
// 		ID:      1,
// 		Name:    "Task One",
// 		Content: "Some Content",
// 	},
// }

// func IndexRouter(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>WELCOME TO MY API</h1>")
// }

// func getTasks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// }

// func createTasks(w http.ResponseWriter, r *http.Request) {
// 	var newTask taskSchema
// 	reqBody, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		fmt.Fprintf(w, "INSER VALID TASK")
// 	}

// 	json.Unmarshal(reqBody, &newTask)
// 	newTask.ID = len(tasks) + 1
// 	tasks = append(tasks, newTask)

// 	w.Header().Set("Content-Type", "application/json") //cabecera
// 	w.WriteHeader(http.StatusCreated)                  //Si el tipo de dato fue correcto
// 	json.NewEncoder(w).Encode(newTask)                 //contendido
// }

// func getTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)["id"]

// 	taskId, err := strconv.Atoi(vars)

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 		return
// 	}

// 	for _, task := range tasks {
// 		if task.ID == taskId {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(task)
// 		}
// 	}

// }

// func deleteTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)["id"]

// 	taskId, err := strconv.Atoi(vars)

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 		return
// 	}

// 	for index, task := range tasks {
// 		if task.ID == taskId {
// 			tasks = append(tasks[:index], tasks[index+1:]...)
// 			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskId)
// 		}
// 	}

// }

// func updateTask(w http.ResponseWriter, r *http.Request) {
// 	var updatedTask taskSchema

// 	vars := mux.Vars(r)["id"]

// 	taskId, err := strconv.Atoi(vars)

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 		return
// 	}

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Please Enter Valid Data")
// 	}
// 	json.Unmarshal(reqBody, &updatedTask)

// 	for i, t := range tasks {
// 		if t.ID == taskId {
// 			tasks = append(tasks[:i], tasks[i+1:]...)

// 			updatedTask.ID = t.ID
// 			tasks = append(tasks, updatedTask)
// 			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskId)
// 		}
// 	}
// }

// func main() {
// 	r := mux.NewRouter().StrictSlash(true)
// 	r.HandleFunc("/", IndexRouter)
// 	r.HandleFunc("/task", getTasks).Methods("GET")
// 	r.HandleFunc("/task", createTasks).Methods("POST")
// 	r.HandleFunc("/task/{id}", getTask).Methods("GET")
// 	r.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
// 	r.HandleFunc("/task/{id}", updateTask).Methods("PUT")

// 	log.Fatal(http.ListenAndServe(":3005", r))
// }
//