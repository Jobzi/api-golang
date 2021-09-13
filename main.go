package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserSchema struct {
	id       string `json:"id"`
	user     string `json:"user"`
	password string `json:"user"`
}

type Users []UserSchema

var users = Users{
	{
		id:       "Qwerty",
		user:     "hola@gmail.com",
		password: "Hola",
	},
}

func IndexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, users)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexRouter)

	http.ListenAndServe(":3005", r)
}
