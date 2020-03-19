package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

const url  ="http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,form map[string] string) string
}

func download(r Retriever) string{
	return  r.Get("http://www.imooc.com")
}

func post(poster Poster){
	poster.Post("http://www/imooc.com",
		map[string] string{
			"name":"ccmoue",
			"course":"golang",
	})
}
//接口组合
type RetriverPoster interface{
	Retriever
	Poster
}
//调用接口组合
func session(s RetriverPoster) string{
	 s.Post(url,map[string]string{
	 	"contents":"another fake imooc",
	 })
	 return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is fake imooc.com"}
	r =&retriever
	//fmt.Printf("%T %V\n",r,r)
	r =real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	//指针接收 要引用的话要加&
	//r1 = &real.Retriever{
	//	UserAgent:"Mozilla/5.0",
	//	TimeOut:time.Minute,}

	//fmt.Printf("%T %V\n",r,r)
	//fmt.Println(download(r))
	//fmt.Println(download(r1))

	inspect(r)
	//type assertion
	realRetriever :=r.(real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	fmt.Println("try a session")

	fmt.Println(session(&retriever))
}


func inspect(r Retriever){
	fmt.Printf("%T %V\n",r,r)
	switch v:=r.(type){
	case *mock.Retriever:
		fmt.Println("Contents",v.Contents)
	case real.Retriever:
		fmt.Println("userAgent",v.UserAgent)
	}
}
