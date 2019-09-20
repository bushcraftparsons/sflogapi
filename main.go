package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sflogapi/app"
	"sflogapi/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//Curl test
/*
 curl -H "Origin: http://localhost:3001" \
 -H "Access-Control-Request-Method: POST" \
 -H "Access-Control-Request-Headers: X-Requested-With, Authorization" \
 -X OPTIONS --verbose http://localhost:8080/login
*/
func main() {
	//Loads the .env file with db connection info
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	//Update security group for DB with current ip if running on dev
	//https://console.aws.amazon.com/ec2/v2/home?region=us-east-1#SecurityGroups:search=sg-1cdc4442;sort=groupId

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/", homeLink).Methods("GET")
	// router.HandleFunc("/login", controllers.Authenticate).Methods("POST, OPTIONS, PUT, HEAD, GET")
	router.HandleFunc("/login", controllers.Authenticate).Methods("POST")

	port := os.Getenv("HOSTPORT") //Get port, not set locally, needs to be set on production
	if port == "" {
		port = "8080" //localhost
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Origin", "Referer", "Sec-Fetch-Mode", "User-Agent"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router))) //Launch the app, visit localhost:8080/api
}
