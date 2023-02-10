// user.go

package models

import "gorm.io/gorm"

// gorm.Modelの中身
// type Model struct {
// 	ID			uint `gorm:"primarykey"`
//	CreatedAt	time.Time
// 	UpdatedAt	time.Time
// 	DeletedAt	DeletedAt	`gorm:"index"`
//}

type User struct {
	gorm.Model
	FirstName 	string
	LastName		string
	Email 			string	`gorm:"unique"`
	Password		[]byte
}