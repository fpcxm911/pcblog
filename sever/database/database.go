package database

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID         int64  `gorm:"column:user_id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

type Comment struct {
	ID              int64  `gorm:"column:comment_id"`
	UserID          int64  `gorm:"column:user_id"`
	ArticleID       int64  `gorm:"column:article_id"`
	CommentDate     int64  `gorm:"column:comment_date"`
	CommentContent  string `gorm:"column:comment_content"`
	ParentCommentID int64  `gorm:"column:parent_comment_id"`
	LikeCnt         int64  `gorm:"column:comment_like_count"`
}

func (c Comment) TableName() string {
	return "comments"
}

type Follower struct {
	IdentifyID   int64 `gorm:"column:id"`
	UserID       int64 `gorm:"column:user_id"`
	UserFollower int64 `gorm:"column:user_follower"`
}

func (f Follower) TableName() string {
	return "user_follower"
}

type Subscriber struct {
	IdentifyID     int64 `gorm:"column:id"`
	UserID         int64 `gorm:"column:user_id"`
	UserSubscriber int64 `gorm:"column:user_subscriber"`
}

func (s Subscriber) TableName() string {
	return "user_subscriber"
}

type Article struct {
	ArticleID      int64  `gorm:"column:article_id"`
	UserID         int64  `gorm:"column:user_id"`
	ArticleTitle   string `gorm:"column:article_title"`
	ArticleContent string `gorm:"column:article_content"`
	ArticleDate    int64  `gorm:"column:article_date"`
}

func (a Article) TableName() string {
	return "article"
}

type Inbox struct {
	ID             int64  `gorm:"column:id"`
	ArticleID      int64  `gorm:"column:article_id"`
	UserID         int64  `gorm:"column:user_id"`
	ArticleTitle   string `gorm:"column:article_title"`
	ArticleContent string `gorm:"column:article_content"`
	ArticleDate    int64  `gorm:"column:article_date"`
}

func (i Inbox) TableName() string {
	return "inbox"
}

func SetupDefaultDatabase() *gorm.DB {
	username := "root"
	password := "19990918pm12"
	host := "127.0.0.1"
	port := 3306
	Dbname := "blog"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
	}

	return db
}
