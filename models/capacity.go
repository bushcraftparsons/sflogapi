package models

import (
	u "sflogapi/utils"
)

//Capacity is a struct to rep Capacity record Primary key (userid, Capacity, type)
type Capacity struct {
	UserID   int    `gorm:"primary_key; column:user_id" json:"userid,omitempty"`
	Capacity string `gorm:"primary_key; column:capacity" json:"capacity,omitempty"`
}

//ListCapacity returns a list of all the Capacity
func ListCapacity(userid int) map[string]interface{} {
	Capacity := []Capacity{}
	err := GetDB().Table("capacity").Where("user_id = ?", userid).Find(&Capacity)
	if err.Error != nil {
		return u.Message(false, "Error getting list of Capacity")
	}
	resp := u.Message(true, "Capacity returned")
	resp["capacity"] = Capacity
	return resp
}

//AddCapacity adds a new Capacity to the db
func AddCapacity(newCapacity Capacity) map[string]interface{} {
	//Assume if there is no error between adding and returning a message that this has succeeded.
	//Supposed to use newRecord, but this doesn't work. Possibly due to composite primary key
	GetDB().Table("capacity").Create(&newCapacity)
	return u.Message(true, "Capacity added: "+newCapacity.Capacity)
}

//DropCapacity deletes a Capacity from the db
func DropCapacity(dropCapacity Capacity) map[string]interface{} {
	//Check primary key is in Capacity.
	if dropCapacity.Capacity != "" && dropCapacity.UserID != 0 {
		//primary key exists, go ahead and delete
		GetDB().Table("capacity").Delete(&dropCapacity)
		return u.Message(true, "Capacity deleted")
	} else {
		return u.Message(false, "Primary key didn't exist to delete Capacity")
	}
}
