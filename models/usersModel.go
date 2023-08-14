package models

import "time"

type User struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	DatePassword time.Time
	DateVerify   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
