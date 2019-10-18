package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sflogapi/models"
	u "sflogapi/utils"
)

//ListAircraft returns a list of aircraft
var ListAircraft = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	resp := models.ListAircraft(userid)
	u.Respond(w, resp)
}

//AddAircraft adds a new aircraft to the db
var AddAircraft = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	newCraft := &models.Aircraft{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &newCraft) //decode the request body into struct and fail if any error occurs
	if err != nil {
		fmt.Println("Error", err)
		u.Respond(w, u.Message(false, "Failed decoding to aircraft struct"))
		return
	}

	//Now add userid to struct
	newCraft.UserID = userid

	resp := models.AddAircraft(*newCraft)
	u.Respond(w, resp)
}

//DeleteAircraft deletes the aircraft from the database
var DeleteAircraft = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	dropCraft := &models.Aircraft{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &dropCraft) //decode the request body into struct and fail if any error occurs
	if err != nil {
		fmt.Println("Error", err)
		u.Respond(w, u.Message(false, "Failed decoding to aircraft struct"))
		return
	}

	//Now add userid to struct
	dropCraft.UserID = userid

	resp := models.DropAircraft(*dropCraft)
	u.Respond(w, resp)
}
