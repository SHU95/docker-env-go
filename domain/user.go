package domain

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Users []User
