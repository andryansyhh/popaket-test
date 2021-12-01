package entity

type Logistics struct {
	LogisticName    string `json:"logistic_name"`
	Amount          int    `json:"amount"`
	DestinationName string `json:"destination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}

type LogisticInput struct {
	LogisticName    string `json:"logistic_name"`
	Amount          int    `json:"amount"`
	DestinationName string `json:"destination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}
