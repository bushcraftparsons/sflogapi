package models

import (
	u "sflogapi/utils"
)

//Aircraft is a struct to rep aircraft record Primary key (userid, aircraft, type)
type Aircraft struct {
	UserID   int    `gorm:"primary_key; column:user_id" json:"userid,omitempty"`
	Aircraft string `gorm:"primary_key; column:aircraft" json:"aircraft,omitempty"`
	Type     string `gorm:"primary_key; column:type" json:"type,omitempty"`
}

//ListAircraft returns a list of all the aircraft with types
func ListAircraft(userid int) map[string]interface{} {
	aircraft := []Aircraft{}
	err := GetDB().Table("aircraft").Where("user_id = ?", userid).Find(&aircraft)
	if err.Error != nil {
		return u.Message(false, "Error getting list of aircraft")
	}
	resp := u.Message(true, "Aircraft returned")
	resp["aircraft"] = aircraft
	return resp
}

//AddAircraft adds a new aircraft to the db
func AddAircraft(newCraft Aircraft) map[string]interface{} {
	//Assume if there is no error between adding and returning a message that this has succeeded.
	//Supposed to use newRecord, but this doesn't work. Possibly due to composite primary key
	GetDB().Table("aircraft").Create(&newCraft)
	return u.Message(true, "Aircraft added: "+newCraft.Aircraft+"|"+newCraft.Type)
}

//DropAircraft deletes an aircraft from the db
func DropAircraft(dropCraft Aircraft) map[string]interface{} {
	//Check primary key is in aircraft.
	if dropCraft.Aircraft != "" && dropCraft.Type != "" && dropCraft.UserID != 0 {
		//primary key exists, go ahead and delete
		GetDB().Table("aircraft").Delete(&dropCraft)
		return u.Message(true, "Aircraft deleted")
	} else {
		return u.Message(false, "Primary key didn't exist to delete aircraft")
	}
}
