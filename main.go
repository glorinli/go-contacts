package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/glorinli/go-contacts/controllers"

	"github.com/glorinli/go-contacts/app"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api", controllers.ListApi).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/me/contacts/new", controllers.CreateContact).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Port is:", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print("Fail to start server", err)
	}
}
