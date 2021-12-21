package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	blog "pcblog/blog"
)

type BlogService struct {
	blog.UserService
	blog.ArticleService
}

func main() {
	err := blog.RegisterBlogService(new(BlogService))
	if err != nil {
		fmt.Println(err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
