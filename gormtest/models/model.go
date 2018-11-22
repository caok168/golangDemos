package models

import (
	"encoding/json"
	"time"
)

// Model represents a base model with primary id key and timestamps
type Model struct {
	ID int `gorm:"primary_key" json:"id" example:"1"`

	CreatedBy  int        `json:"createdBy"`
	CreatedAt  time.Time  `json:"createdAt" example:"2018-07-07T13:38:13+08:00"`
	ModifiedAt time.Time  `json:"modifiedAt" example:"2018-07-07T13:38:13+08:00"`
	DeletedAt  *time.Time `json:"-" example:"2018-07-07T13:38:13+08:00"`

	Extras *json.RawMessage `json:"extras,omitempty"` // reserved
}
