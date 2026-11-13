package models

import "time"

type User struct{
	ID 			uint 		`gorm:"primaryKey" json:"id"`
	Name 		string 		`gorm:"size:100;not null" json:"name"`
	Email 		string 		`gorm:"size:255; uniqueIndex;not null" json:"email"`
	Role 		string 		`gorm:"size:20;default:user;not null" json:"role"`
	Password 	string 		`gorm:"not null" json:"-"`
	CreateAt 	time.Time   `json:"create_at"`
}