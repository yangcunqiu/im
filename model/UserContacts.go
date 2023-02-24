package model

import "gorm.io/gorm"

type UserContacts struct {
	gorm.Model

	UserId         uint
	ContactsUserId uint
}

func (userContacts UserContacts) TableName() string {
	return "user_contacts"
}
