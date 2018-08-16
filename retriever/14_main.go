package main

import (
	"fmt"
	"learn-go/retriever/mock"
	real2 "learn-go/retriever/real"
	"time"
)

const url = "https://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

// 组合两个接口到一个接口中
type PosterRetriever interface {
	Retriever
	Poster
}

func session(s PosterRetriever) string {
	s.Post(url, map[string]string{
		"content": "another faked imooc.com",
	})

	return s.Get(url)
}

func inspect(r Retriever) {

	fmt.Println(r)
	fmt.Printf(">	%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Content:", v.Content)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {

	var r Retriever
	//fmt.Println(download(r))

	r = &mock.Retriever{"this is a fake imooc.com"}
	fmt.Println(download(r))
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	// 查看接口的两种方式：1.Type switch 2.Type assertion
	inspect(r)
	//fmt.Println(download(r2))

	fmt.Println("-----------------------------------")

	// Type assertion
	if retriver1, ok := r.(*real2.Retriever); ok {
		fmt.Println(retriver1.UserAgent)
	}
	if retriver2, ok := r.(*mock.Retriever); ok {
		fmt.Println(retriver2.Content)
	}

	fmt.Println("-----------------------------------")
	pr := &mock.Retriever{"this is a fake imooc.com"}
	fmt.Println(session(pr))
}
