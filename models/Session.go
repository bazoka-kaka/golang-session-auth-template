package models

import "time"

type Session struct {
	Username string
	Expiry   time.Time
	Value    string
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
