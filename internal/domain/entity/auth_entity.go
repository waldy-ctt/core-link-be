package entity

import "time"

type Auth struct {
	UserId       string    `json:"userId"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
}
