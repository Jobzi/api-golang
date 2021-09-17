package main

import (
	t "API/controller"
	mi "API/middlewares"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>WELCOME TO MY API</h1>")
}


func main() {
	r := mux.NewRouter().StrictSlash(true)
	
	//middleware
	r.Use(mi.LoggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	
	//endpoints
	r.HandleFunc("/", IndexRouter)
	r.HandleFunc("/task", t.GetAllTask).Methods("GET")
	r.HandleFunc("/task", t.PostTask).Methods("POST")
	
	//server
	log.Fatal(http.ListenAndServe(":3005", r))
}

//this method list all databases of proyect
// databases, err := client.ListDatabaseNames(ctx, bson.M{})
// fmt.Println(databases)