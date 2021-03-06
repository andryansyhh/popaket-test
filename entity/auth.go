package entity

import (
	"time"
)

type Auths struct {
	ID        string    `gorm:"PrimaryKey" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Msisdn    string    `json:"msisdn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthInputs struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Msisdn   string `json:"msisdn" binding:"required"`
}

type LoginAuthInputs struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
