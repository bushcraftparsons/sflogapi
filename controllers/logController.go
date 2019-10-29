package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sflogapi/models"
	u "sflogapi/utils"
	"time"
)

type indexData struct {
	start int
	end   int
}

//AddLog sends the request to be added to the logs
var AddLog = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	log := &models.Log{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.Unmarshal(b, &log) //decode the request body into struct and fail if any error occurs
	if err != nil {
		fmt.Println("Error", err)
		u.Respond(w, u.Message(false, "Failed decoding to log struct"))
		return
	}

	//Now add userid to struct
	log.UserID = userid
	//And time stamps
	log.CreatedAt = time.Now()
	log.UpdatedAt = time.Now()

	resp := models.AddLog(*log)
	u.Respond(w, resp)
}

/*
ListLogs lists all the logs for the given user between the given indices
*/
var ListLogs = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	decoder := json.NewDecoder(r.Body)

	var data indexData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	resp := models.ListLogs(userid, data.start, data.end)
	u.Respond(w, resp)
}
