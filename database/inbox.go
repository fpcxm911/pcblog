package database

import "gorm.io/gorm"

func (i Inbox) FindInbox(db *gorm.DB, user User) (inboxes []Inbox, err error) {
	err = db.Where("user_id = ?", user.ID).Order("article_date").Find(&inboxes).Error

	return
}

func (i Inbox) AddInbox(db *gorm.DB) (err error) {
	err = db.Create(&i).Error
	return
}

func (i Inbox) RemoveInbox(db *gorm.DB) (err error) {
	err = db.Delete(&i).Error
	return
}

func (i Inbox) SubUpdate(db *gorm.DB, user User, author User) (err error) {
	var a Article
	articles, err := a.Find(db, author)
	if err != nil {
		return err
	}
	for _, article := range articles {
		i = Inbox{
			ArticleID:      article.ArticleID,
			UserID:         user.ID,
			AuthorUserID:   author.ID,
			ArticleTitle:   article.ArticleTitle,
			ArticleContent: article.ArticleContent,
			ArticleDate:    article.ArticleDate,
		}
		err := i.AddInbox(db)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i Inbox) UnsubUpdate(db *gorm.DB, user User, author User) (err error) {
	var deletedArticles []Article
	err = db.Where("user_id = ? AND author_id = ?", user.ID, author.ID).Find(&deletedArticles).Error
	if err != nil {
		return
	}
	err = db.Delete(&deletedArticles).Error
	return
}
