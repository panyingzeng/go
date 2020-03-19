package main

import (
	"fmt"
	"time"
)

func worker(c chan  int){
	for{
		n := <-c //收channel数字
		fmt.Println(n)
	}
}

func workerx(id int , c chan  int){
	for n :=range c {
		/*n, ok :=<-c
		if !ok {
			break
		}*/
		fmt.Printf("Worker %d Received %c\n",id,n)
	}
}

//返回一个channel
func CreateWorker(id int ) chan  int{
	c :=make(chan  int)
	go workerx(id, c)
	return c
}
func chanDemo(){
	//var c  chan  int // c==nil
	var channels[10] chan int
	for i:=0;i<10;i++{
		channels[i] =make(chan int)
		//go worker(c)
		channels[i] =CreateWorker(i)
	}

	for i :=0;i<10;i++{
		channels[i] <-'a'+i
	}
	for i :=0;i<10;i++{
		channels[i] <-'A'+i
	}
	//c <-1 //发channel数字 chan <- int

	time.Sleep(time.Millisecond)
}

func bufferedChannel(){
	c :=make(chan int, 3)
	go workerx(0,c)
	c<-'1'
	c<-'2'
	c<-'3'
	c<-'4'
	time.Sleep(time.Millisecond)
}

func channelClose(){
	c :=make(chan int, 3)
	go workerx(0,c)
	c<-'1'
	c<-'2'
	c<-'3'
	c<-'4'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	fmt.Println("channel as first-class citieszon")
	chanDemo()
	fmt.Println("bufferedChannel")
	bufferedChannel()
	fmt.Println("channelclose in range")
	channelClose()//发送方close（）
}
