package models

import "time"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	Username string
	Expiry   time.Time
	Value    string
}
