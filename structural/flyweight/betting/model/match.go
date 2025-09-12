package model

import "time"

type Match struct {
	Date time.Time
	VisitorID uint64
	LocalID uint64
	LocalScore uint8
	VisitorScore uint8
	LocalShoots uint8
	VisitorShoots uint8
}