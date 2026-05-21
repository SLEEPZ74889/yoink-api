package model

import "time"

type Click struct {
	ID        string
	LinkID    string
	ClickedAt time.Time
	Referrer  string
	Browser   string
	OS        string
	Device    string
	IPHash    string
}
