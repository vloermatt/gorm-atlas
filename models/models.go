package models

import "time"

type User struct {
	CreatedAt time.Time // Automatically managed by GORM at create time
	UpdatedAt time.Time // Automatically managed by GORM at updated time
	Name      string
	Surname   string
	ID        string `gorm:"primary_key,autoIncrement:false"`
}
