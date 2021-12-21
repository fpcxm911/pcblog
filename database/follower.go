package database

import (
	"errors"

	"gorm.io/gorm"
)

func (f Follower) Update(db *gorm.DB) (err error) {
	var result Follower
	err = db.Where("user_id = ? AND user_follower = ?", f.UserID, f.UserFollower).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// result not found which means user wants to sub
		if err = db.Create(&f).Error; err != nil {
			return
		}
		return nil
	}

	// result is found, which mean user want to unsub
	if err = db.Delete(&result).Error; err != nil {
		return
	}

	return nil
}

func (f Follower) Find(db *gorm.DB, authorId int64) (results []Follower, err error) {
	err = db.Where("user_id = ?", authorId).Find(&results).Error
	return
}
