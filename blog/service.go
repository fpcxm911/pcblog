package pcblog

import (
	"net/rpc"
	"pcblog/database"

	"gorm.io/gorm"
)

// ServiceName blog service
const ServiceName = "BlogService"
const LoginSuccess = "New account Registered. Login completely. Here is the inbox."
const RegisterUsernameUsed = "this username has been taken. Please try a new one"
const WrongUsername = "the provided username doesn't exist, please register first"
const WrongPassword = "wrong password"

// UserInterface User related interface
type UserInterface interface {
	//Register sign up a new account
	Register(request database.User, reply *LoginResponse) error
	//Login login into an account
	Login(request database.User, reply *LoginResponse) error
	//Subscribe subscribe another user
	Subscribe(request [2]database.User, reply *string) error
	//Unsubscribe unsub a user
	Unsubscribe(request [2]database.User, reply *string) error
}

// ArticleInterface Article related methods
type ArticleInterface interface {
	//NewArticle publish new article
	NewArticle(request database.ArticleToPublish, reply *string) error
}

// InboxInterface method that can perform at inbox
//type InboxInterface interface {
//	//Subscribe subscribe another user
//	Subscribe(request []database.User, reply *string) error
//	//Unsubscribe unsub a user
//	Unsubscribe(request []database.User, reply *string) error
//}

// ServiceInterface interface for the rpc service, define the functions that blog service provides
type ServiceInterface interface {
	UserInterface
	ArticleInterface
	//InboxInterface
}

// RegisterBlogService register the rpc service on svc
func RegisterBlogService(svc ServiceInterface) error {
	return rpc.RegisterName(ServiceName, svc)
}

// LoginResponse defines the response needed if login or register completely
type LoginResponse struct {
	Db          *gorm.DB
	LoginStatus string
	Inbox       []database.Inbox
	UserInfo    database.User
}
