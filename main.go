package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sflogapi/app"
	"sflogapi/controllers"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
func main() {
	//Update security group for DB with current ip if running on dev
	//https://console.aws.amazon.com/ec2/v2/home?region=us-east-1#SecurityGroups:search=sg-1cdc4442;sort=groupId

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/login", controllers.Authenticate).Methods("POST")

	port := os.Getenv("HOSTPORT") //Get port, not set locally, needs to be set on production
	if port == "" {
		port = "8080" //localhost
	}

	log.Fatal(http.ListenAndServe(":8080", router)) //Launch the app, visit localhost:8000/api
}
