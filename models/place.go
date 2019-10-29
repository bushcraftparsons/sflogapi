package models

import (
	u "sflogapi/utils"
)

//Place is a struct to rep place record Primary key (userid, Place, type)
type Place struct {
	UserID int    `gorm:"primary_key; column:user_id" json:"userid,omitempty"`
	Place  string `gorm:"primary_key; column:place" json:"place,omitempty"`
}

//ListPlaces returns a list of all the Places
func ListPlaces(userid int) map[string]interface{} {
	Places := []Place{}
	err := GetDB().Where("user_id = ?", userid).Find(&Places)
	if err.Error != nil {
		return u.Message(false, "Error getting list of places")
	}
	resp := u.Message(true, "places returned")
	resp["places"] = Places
	return resp
}

//AddPlace adds a new Place to the db
func AddPlace(newPlace Place) map[string]interface{} {
	//Assume if there is no error between adding and returning a message that this has succeeded.
	//Supposed to use newRecord, but this doesn't work. Possibly due to composite primary key
	GetDB().Create(&newPlace)
	return u.Message(true, "place added: "+newPlace.Place)
}

//DropPlace deletes a Place from the db
func DropPlace(dropPlace Place) map[string]interface{} {
	//Check primary key is in Place.
	if dropPlace.Place != "" && dropPlace.UserID != 0 {
		//primary key exists, go ahead and delete
		GetDB().Delete(&dropPlace)
		return u.Message(true, "place deleted")
	} else {
		return u.Message(false, "Primary key didn't exist to delete place")
	}
}
