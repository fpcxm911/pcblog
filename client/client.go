package main

import (
	"fmt"
	"log"
	"net/rpc"
	blog "pcblog/blog"
	"pcblog/database"
)

// BlogClient is the implementation of ServiceInterface from the client side
type BlogClient struct {
	*rpc.Client
}

var _ blog.ServiceInterface = (*BlogClient)(nil)

func (b *BlogClient) Register(request database.User, reply *blog.LoginResponse) error {
	return b.Client.Call(blog.ServiceName+".Register", request, reply)
}

func (b *BlogClient) Login(request database.User, reply *blog.LoginResponse) error {
	return b.Client.Call(blog.ServiceName+".Login", request, reply)
}

func (b *BlogClient) Subscribe(request [2]database.User, reply *string) error {
	return b.Client.Call(blog.ServiceName+".Subscribe", request, reply)
}

func (b *BlogClient) Unsubscribe(request [2]database.User, reply *string) error {
	return b.Client.Call(blog.ServiceName+".Unsubscribe", request, reply)
}

func (b *BlogClient) NewArticle(request database.ArticleToPublish, reply *string) error {
	return b.Client.Call(blog.ServiceName+".NewArticle", request, reply)
}

var _ blog.ServiceInterface = (*BlogClient)(nil)

func DialBlogService(network, address string) (*BlogClient, error) {
	c, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &BlogClient{Client: c}, nil
}

func main() {
	client, err := DialBlogService("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Err Dial Client:", err)
	}

	//test register new account
	clientTestUser := database.User{
		Username: "clientTest3",
		Password: "admin",
	}
	var loginResponse blog.LoginResponse
	err = client.Register(clientTestUser, &loginResponse)
	if err != nil {
		log.Println(err)
	}
	log.Println(loginResponse)

	//test subscribe
	clientTestUser1 := database.User{
		Username: "clientTest1",
		Password: "admin",
	}
	clientTestUser2 := database.User{
		Username: "clientTest2",
		Password: "admin",
	}
	clientTestUser3 := database.User{
		Username: "clientTest3",
		Password: "admin",
	}
	request := [2]database.User{clientTestUser3, clientTestUser1}
	var reply *string
	err = client.Unsubscribe(request, reply)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("User 2 sub user 1 successfully!")

	// test publish new article
	var reply string
	article := database.ArticleToPublish{
		UserID:         5,
		ArticleTitle:   "TestUser1's article",
		ArticleContent: "I want to see if I can receive rpc reply",
		ArticleDate:    0,
	}
	err = client.NewArticle(article, &reply)
	fmt.Println(reply)
}
