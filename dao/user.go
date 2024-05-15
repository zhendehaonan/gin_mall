package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// 创建用户
func (userDao *UserDao) CreateUser(user *model.User) error {
	return userDao.DB.Model(&model.User{}).Create(&user).Error
}

// 根据id查找用户信息
func (userDao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = userDao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// 根据id修改信息
func (userDao *UserDao) UpdateUserById(id uint, user *model.User) error {
	return userDao.DB.Model(&model.User{}).Where("id=?", id).Updates(&user).Error
}
