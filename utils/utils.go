package utils

import (
	"encoding/json"
	"net/http"
)

//Message given a status boolean and message string, returns a map of the status and message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond json encodes the data and adds a json header
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
