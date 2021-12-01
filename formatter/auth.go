package formatter

import (
	"popaket/entity"
	"time"
)

type UserFormat struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Msisdn   string `json:"msisdn"`
}

type UserDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.Auths) UserFormat {
	var formatUser = UserFormat{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Msisdn:   user.Msisdn,
	}

	return formatUser
}


// func FormatDeleteUser(msg string) UserDeleteFormat {
// 	var deleteFormat = UserDeleteFormat{
// 		Message:    msg,
// 		TimeDelete: time.Now(),
// 	}

// 	return deleteFormat
// }
