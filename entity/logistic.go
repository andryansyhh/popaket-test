package entity

import (
	"time"
)

type Logistics struct {
	ID              string    `gorm:"PrimaryKey" json:"id"`
	LogisticName    string    `json:"logistic_name"`
	Amount          int       `json:"amount"`
	DestinationName string    `json:"destination_name"`
	OriginName      string    `json:"origin_name"`
	Duration        string    `json:"duration"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type LogisticInput struct {
	LogisticName    string `json:"logistic_name"`
	Amount          int    `json:"amount"`
	DestinationName string `json:"destination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}
