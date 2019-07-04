package main

import (
	"github.com/gorilla/mux"
	"go-contacts/app"
	"os"
	"fmt"
	"net/http"
)

func main() {
	router := mux.NewRounter()
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Port is:", port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print("Fail to start server", err)
	}
}