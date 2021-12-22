package database

import "gorm.io/gorm"

func (a *Article) Find(db *gorm.DB, author User) (articles []Article, err error) {
	err = db.Where("user_id = ?", author.ID).Order("article_date desc").Find(&articles).Error
	if err != nil {
		return
	}
	return
}

func (a *Article) Add(db *gorm.DB) (err error) {
	err = db.Create(&a).Error
	return
}
