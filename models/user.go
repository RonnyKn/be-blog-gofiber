package models

type User struct {
	UID      uint   `json:"uid" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null;column:username;size:255"`
	Password string `json:"password" gorm:"not null;column:password;size:50"`
	Email    string `json:"email" gorm:"not null;column:email;size:100"`
	BlogID   uint   `json:"blogid" gorm:"foreignKey:CompanyRefer"`
}