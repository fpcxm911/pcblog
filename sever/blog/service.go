package blog

import (
	"github.com/fpcxm911/pcblog/sever/blog/database"
	"net/rpc"
)

// ServiceName
const ServiceName = "BlogService"

// UserInterface User related interface
type UserInterface interface {
	//Register sign up a new account
	Register(request AccountInfo, reply *string) error
	//Login login into an account
	Login(request AccountInfo, reply *string) error
	//Subscribe subscribe another user
	Subscribe(request []AccountInfo, reply *string) error
	//Unsubscribe unsub a user
	Unsubscribe(request []AccountInfo, reply *string) error
}

// ArticleInterface Article related interface
type ArticleInterface interface {
	//NewArticle publish new article
	NewArticle(request Article, reply *string) error
}

// ServiceInterface interface for the rpc service, define the functions that blog service provides
type ServiceInterface interface {
	UserInterface
	ArticleInterface
}

// RegisterBlogService register the rpc service on svc
func RegisterBlogService(svc ServiceInterface) error {
	return rpc.RegisterName(ServiceName, svc)
}

// AccountInfo Defines the info of an account
type AccountInfo struct {
	databse.Account
}

// Article Defines the element of an article
type Article struct {
}
