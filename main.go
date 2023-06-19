package main

import (
	"log"
	"mvc/controllers"
	"net/http"
)

func main() {

	http.HandleFunc("/", controllers.HomeRoute)
	http.HandleFunc("/login", controllers.LoginRoute)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}
