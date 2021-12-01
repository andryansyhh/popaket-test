package migration

import "time"

type Users struct {
	ID        string `gorm:"PrimaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	Name      string
	Msisdn    string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Logistics struct {
	ID              string `gorm:"PrimaryKey"`
	LogisticName    string
	Amount          int
	DestinationName string
	OriginName      string
	Duration        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
