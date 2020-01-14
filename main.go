package main

import (
	"log"
	"net/http"
	"sflogapi/app"
	"sflogapi/controllers"

	u "sflogapi/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//Package autocert provides automatic access to certificates from Let's Encrypt and any other ACME-based CA.
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
	router.HandleFunc("/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/addLog", controllers.AddLog).Methods("POST")
	router.HandleFunc("/listLogs", controllers.ListLogs).Methods("POST")

	router.HandleFunc("/listAircraft", controllers.ListAircraft).Methods("GET")
	router.HandleFunc("/addAircraft", controllers.AddAircraft).Methods("POST")
	router.HandleFunc("/deleteAircraft", controllers.DeleteAircraft).Methods("POST")

	router.HandleFunc("/listPlaces", controllers.ListPlaces).Methods("GET")
	router.HandleFunc("/addPlace", controllers.AddPlace).Methods("POST")
	router.HandleFunc("/deletePlace", controllers.DeletePlace).Methods("POST")

	router.HandleFunc("/listCapacity", controllers.ListCapacity).Methods("GET")
	router.HandleFunc("/addCapacity", controllers.AddCapacity).Methods("POST")
	router.HandleFunc("/deleteCapacity", controllers.DeleteCapacity).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Origin", "Referer", "Sec-Fetch-Mode", "User-Agent"})
	originsOk := handlers.AllowedOrigins([]string{"*", "http://localhost:3001", "https://sfl.formyer.com", "http://86.184.53.15:3001"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	// using TLS
	// Includes redirect from http above to https
	log.Fatal(http.ListenAndServeTLS(":8080", "apiserver.crt", "apiserver.key", handlers.CORS(originsOk, headersOk, methodsOk)(router))) //Launch the app, visit https://localhost:8080/api
}
