package storage

import (
	"popaket/entity"

	"gorm.io/gorm"
)

type UserDao interface {
	RegisterUser(user entity.Users) (entity.Users, error)
	FindUserByUsername(username string) (entity.Users, error)
}

type dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) RegisterUser(user entity.Users) (entity.Users, error) {

	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *dao) FindUserByUsername(username string) (entity.Users, error) {
	var user entity.Users

	if err := r.db.Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil

}
