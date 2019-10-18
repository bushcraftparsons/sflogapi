package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sflogapi/models"
	u "sflogapi/utils"
)

//ListCapacity returns a list of Capacity
var ListCapacity = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	resp := models.ListCapacity(userid)
	u.Respond(w, resp)
}

//AddCapacity adds a new Capacity to the db
var AddCapacity = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	newCapacity := &models.Capacity{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &newCapacity) //decode the request body into struct and fail if any error occurs
	if err != nil {
		fmt.Println("Error", err)
		u.Respond(w, u.Message(false, "Failed decoding to Capacity struct"))
		return
	}

	//Now add userid to struct
	newCapacity.UserID = userid

	resp := models.AddCapacity(*newCapacity)
	u.Respond(w, resp)
}

//DeleteCapacity deletes the Capacity from the database
var DeleteCapacity = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	dropCapacity := &models.Capacity{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &dropCapacity) //decode the request body into struct and fail if any error occurs
	if err != nil {
		fmt.Println("Error", err)
		u.Respond(w, u.Message(false, "Failed decoding to Capacity struct"))
		return
	}

	//Now add userid to struct
	dropCapacity.UserID = userid

	resp := models.DropCapacity(*dropCapacity)
	u.Respond(w, resp)
}
