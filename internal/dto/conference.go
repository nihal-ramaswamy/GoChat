package dto

import "time"

type Conference struct {
	Code      string    `json:"code"`
	Admin     string    `json:"admin"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

func NewConference(code, admin string) *Conference {
	return &Conference{
		Code:   code,
		Admin:  admin,
		Active: true,
	}
}
