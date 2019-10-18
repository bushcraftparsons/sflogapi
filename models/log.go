package models

import (
	"fmt"
	u "sflogapi/utils"
	"time"

	"github.com/jinzhu/gorm"
)

//Log is a struct to rep log record
type Log struct {
	gorm.Model
	ID             int       `gorm:"PRIMARY_KEY"`
	UserID         int       `gorm:"user_id" json:"userid,omitempty"`
	Date           time.Time `gorm:"date" json:"date,omitempty"`
	Aircraft       string    `gorm:"aircraft" json:"aircraft,omitempty"`
	Type           string    `gorm:"type" json:"type,omitempty"`
	DeparturePlace string    `gorm:"departure_place" json:"depPlace,omitempty"`
	DepartureTime  time.Time `gorm:"departure_time" json:"depTime,omitempty"`
	ArrivalPlace   string    `gorm:"arrival_place" json:"arrPlace,omitempty"`
	ArrivalTime    time.Time `gorm:"arrival_time" json:"arrTime,omitempty"`
	FlightTime     int       `gorm:"flight_time" json:"flightDuration,omitempty"`
	InstApp        bool      `gorm:"inst_app" json:"instrumentApproach,omitempty"`
	NightFlight    int       `gorm:"night_flight" json:"nightFlightDuration,omitempty"`
	Log            bool      `gorm:"log" json:"log,omitempty"`
	Comments       string    `gorm:"comments" json:"comments,omitempty"`
	PilotNo        string    `gorm:"pilot_number" json:"pilotNumber,omitempty"`
	Capacity       string    `gorm:"capacity" json:"capacity,omitempty"`
}

//AddLog adds the given log entry to the database
func AddLog(log Log) map[string]interface{} {
	fmt.Println("Adding log: %d "+log.ArrivalPlace+" "+log.ArrivalTime.Format("2006-02-01 15:04:05"), log.UserID)
	err := GetDB().Create(&log).Error
	if err != nil {
		return u.Message(false, "Error saving log entry")
	}
	resp := u.Message(true, "Log saved")
	return resp
}
