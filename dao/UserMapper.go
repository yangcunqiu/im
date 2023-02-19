package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"im/global"
	"im/model"
)

func GetUserList(offset int, limit int, name string) ([]model.User, int64) {
	var total int64
	userList := make([]model.User, 0)
	// 分页一定要加 db.Model, 要不然查不出来
	tx := global.DB.Model(new(model.User))
	if name != "" {
		tx.Where("name like ?", "%"+name+"%")
	}
	tx.Count(&total).Offset(offset).Limit(limit).Find(&userList)
	return userList, total
}

func GetUser(id uint) (*model.User, error) {
	user := model.User{}
	err := global.DB.First(&user, id).Error
	// go的结构体是值类型, 没有查到也会有默认值, 需要检查一下是否是没有查到
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

func CreateUser(tx *gorm.DB, user *model.User) error {
	return tx.Create(user).Error
}

func UpdateUser(user *model.User) {
	global.DB.Model(&user).Updates(user)
}

func GetUserByName(name string) (*model.User, error) {
	user := model.User{}
	err := global.DB.Where("name = ?", name).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}
