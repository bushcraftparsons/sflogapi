package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sflogapi/models"
	u "sflogapi/utils"
	"time"
)

//AddLog sends the request to be added to the logs
var AddLog = func(w http.ResponseWriter, r *http.Request) {
	var userid int
	userid = u.GetContext(w, r, u.UserID).(int)

	log := &models.Log{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&log)
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

	type indexData struct {
		Start int `json:"start,omitempty"`
		End   int `json:"end,omitempty"`
	}

	dec := json.NewDecoder(r.Body)
	var data indexData
	for {
		if err := dec.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	resp := models.ListLogs(userid, data.Start, data.End)
	u.Respond(w, resp)
}
