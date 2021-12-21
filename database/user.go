package database

import (
	"time"

	"gorm.io/gorm"
)

func (u User) FindOne(db *gorm.DB) (response User, err error) {

	err = db.Where("username = ?", u.Username).First(&response).Error
	return
}

func (u User) CreateNewUser(db *gorm.DB) (userinfo User, err error) {
	u.CreateTime = time.Now().Unix()
	if err = db.Create(&u).Error; err != nil {
		return
	}
	userinfo, err = u.FindOne(db)
	return
}
