package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sflogapi/models"
	u "sflogapi/utils"
)

//ListPlaces returns a list of places
var ListPlaces = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	resp := models.ListPlaces(userid)
	u.Respond(w, resp)
}

//AddPlace adds a new place to the db
var AddPlace = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	newPlace := &models.Place{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &newPlace) //decode the request body into struct and fail if any error occurs
	if err != nil {
		fmt.Println("Error", err)
		u.Respond(w, u.Message(false, "Failed decoding to place struct"))
		return
	}

	//Now add userid to struct
	newPlace.UserID = userid

	resp := models.AddPlace(*newPlace)
	u.Respond(w, resp)
}

//DeletePlace deletes the place from the database
var DeletePlace = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	dropPlace := &models.Place{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &dropPlace) //decode the request body into struct and fail if any error occurs
	if err != nil {
		fmt.Println("Error", err)
		u.Respond(w, u.Message(false, "Failed decoding to place struct"))
		return
	}

	//Now add userid to struct
	dropPlace.UserID = userid

	resp := models.DropPlace(*dropPlace)
	u.Respond(w, resp)
}
