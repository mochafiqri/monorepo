package dtos

import "time"

type TokenUser struct {
	Id  string    `json:"id"`
	Nik string    `json:"nik"`
	Exp time.Time `json:"exp"`
}
