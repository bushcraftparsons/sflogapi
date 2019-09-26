package main

import (
	"log"
	"net/http"
	"sflogapi/app"
	"sflogapi/controllers"

	u "sflogapi/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "Welcome to sfl!")
	resp["user"] = "Susannah Parsons"
	u.Respond(w, resp)
}

//Curl test
/*
 curl -H "Origin: http://localhost:3001" \
 -H "Access-Control-Request-Method: POST" \
 -H "Access-Control-Request-Headers: X-Requested-With, Authorization" \
 -X OPTIONS --verbose http://localhost:8080/login
*/
func main() {
	//Update security group for DB with current ip if running on dev
	//https://console.aws.amazon.com/ec2/v2/home?region=us-east-1#SecurityGroups:search=sg-1cdc4442;sort=groupId

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/", homeLink).Methods("GET")
	// router.HandleFunc("/login", controllers.Authenticate).Methods("POST, OPTIONS, PUT, HEAD, GET")
	router.HandleFunc("/login", controllers.Authenticate).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Origin", "Referer", "Sec-Fetch-Mode", "User-Agent"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))) //Launch the app, visit localhost:8080/api
}
