package main

import (
	"log"
	"net/rpc"
	blog "pcblog/blog"
	"pcblog/database"
)

// BlogClient is the implementation of ServiceInterface from the client side
type BlogClient struct {
	*rpc.Client
}

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

func (b *BlogClient) NewArticle(request database.Article, reply *string) error {
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

	// test register new account
	clientTestUser := database.User{
		Username: "clientTest1",
		Password: "admin",
	}
	var loginResponse blog.LoginResponse
	err = client.Register(clientTestUser, &loginResponse)
	if err != nil {
		log.Println(err)
	}
	log.Println(loginResponse.LoginStatus)
}
