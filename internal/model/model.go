package model

import "time"

type Metric struct {
	Id          string    `db:"id"`                       // UUID
	UserId      string    `json:"user_id" db:"user_id"`   // UUID ?
	EventId     int       `json:"event_id" db:"event_id"` // int ?
	EventName   string    `json:"event_name" db:"event_name"`
	LayoutId    int       `json:"layout_id" db:"layout_id"` // int ?
	LayoutName  string    `json:"layout_name" db:"layout_name"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}
