package pcblog

import (
	"errors"
	"pcblog/database"

	"gorm.io/gorm"
)

func (a *ArticleService) NewArticle(request database.ArticleToPublish, reply *string) error {
	db, err := database.SetupDefaultDatabase()
	if err != nil {
		return err
	}
	err = request.Add(db)
	if err != nil {
		return err
	}
	var f database.Follower
	// find all followers of the author
	followers, err := f.Find(db, request.UserID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// This author doesn't have followers, skip the inbox push process
		return nil
	}
	inboxToAdd := database.Inbox{
		ArticleID:      request.ArticleID,
		UserID:         0,
		AuthorUserID:   request.UserID,
		ArticleTitle:   request.ArticleTitle,
		ArticleContent: request.ArticleContent,
		ArticleDate:    request.ArticleDate,
	}
	// for every follower, push the new article to their inbox individually
	for _, follower := range followers {
		inboxToAdd.UserID = follower.UserFollower
		err = inboxToAdd.AddInbox(db)
		if err != nil {
			return err
		}
	}
	*reply = "New article added successfully."
	return nil
}
