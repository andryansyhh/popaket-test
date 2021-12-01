package migration

import "time"

type Auths struct {
	ID        string `gorm:"PrimaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	Name      string
	Msisdn    string `gorm:"unique"`
}

type Logistics struct {
	LogisticName    string
	Amount          int
	DestinationName string
	OriginName      string
	Duration        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
