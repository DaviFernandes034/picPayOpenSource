package model

import "time"

type Transaction struct {
	Amount        float64
	Sender        User
	Receiver      User
	LocaldateTime time.Time
}