package database

import (
	"errors"

	"gorm.io/gorm"
)

func (s Subscriber) Update(db *gorm.DB) (err error) {
	var result Subscriber
	err = db.Where("user_id = ? AND user_subscriber = ?", s.UserID, s.UserSubscriber).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// sub record not found, meaning the user wants to subscriber
		if err = db.Create(&s).Error; err != nil {
			return
		}
		return nil
	}
	// sub record found, meaning the user wants to unsub
	if err = db.Delete(&result).Error; err != nil {
		return
	}
	return nil
}
