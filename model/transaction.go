package model

import "time"

type Transaction struct {
	TransferID 		int
	Amount       	float64
	SenderID        int
	ReceiverID      int
	LocaldateTime 	time.Time
}

type TransactionResponse struct {
	TransferID     int       `json:"TransferID"`
	Amount         float64   `json:"Amount"`
	Sender         User     `json:"Sender"`
	Receiver       User     `json:"Receiver"`
	LocaldateTime  time.Time `json:"LocaldateTime"`
}
