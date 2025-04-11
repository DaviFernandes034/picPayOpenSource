package model

import "time"

type Transaction struct {
	Amount        float64
	Sender        int
	Receiver      int
	LocaldateTime time.Time
}