package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"im/global"
	"im/model"
)

func GetUserList() []model.User {
	userList := make([]model.User, 0)
	global.DB.Find(&userList)
	return userList
}

func GetUser(id uint) (*model.User, error) {
	user := &model.User{}
	err := global.DB.First(user, id).Error
	// go的结构体是值类型, 没有查到也会有默认值, 需要检查一下是否是没有查到
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return user, nil
}

func CreateUser(tx *gorm.DB, user *model.User) error {
	return tx.Create(user).Error
}
