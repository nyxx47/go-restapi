package main

import (
	"fmt"
	"net/http"
	"os"
	"restapi/app"
	"restapi/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contact/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/profile/contact", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
