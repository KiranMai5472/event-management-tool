package models

import (
	"time"
)

type GetEvent struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Location  string    `json:"location"`
	Host      string    `json:"host"`
} // @name GetEvent

type SlotAccept struct {
	EventAccept int `json:event_id`
} // @name SlotAccept

type CreateEvent struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Host     string `json:"host"`
} // @name GetEvent
