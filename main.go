package main

import (
	"fmt"
	"net/http"
	_"os"
	"restapi/app"
	"restapi/controllers"
	"github.com/spf13/viper"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contact/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/profile/contact", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication)

	port := viper.GetString(`server.address`)
	if port == "" {
		port = viper.GetString(`server.address`)
	}

	fmt.Println(port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println(err)
	}
}
