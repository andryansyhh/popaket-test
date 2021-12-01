package formatter

import (
	"popaket/entity"
)

type LogisticFormat struct {
	LogisticName    string `json:"logistic_name"`
	Amount          int    `json:"amount"`
	DestinationName string `json:"destination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}

func FormatLogistic(logistic entity.Logistics) LogisticFormat {
	var formatLogistic = LogisticFormat{
		LogisticName:    logistic.LogisticName,
		Amount:          logistic.Amount,
		DestinationName: logistic.DestinationName,
		OriginName:      logistic.OriginName,
		Duration:        logistic.Duration,
	}

	return formatLogistic
}
