package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int{
	out :=make(chan int)
	go func() {
		i :=0
		for{
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <-i
			i++
		}
	}()
	return out
}


func workerx(id int , c chan  int){
	for n :=range c {
		/*n, ok :=<-c
		if !ok {
			break
		}*/
		fmt.Printf("Worker %d Received %d\n",id,n)
	}
}
func CreateWorker(id int ) chan <- int{
	c :=make(chan  int)
	go workerx(id, c)
	return c
}

func main() {
	var c1,c2 = generator(),generator()
	var work =CreateWorker(0)

	var values  []int
	t :=time.After(10*time.Second)
	tick :=time.Tick(time.Second)
	for  {
		var activeWork chan<- int
		var activeValue int
		if len(values) >0 {
			activeWork = work
			activeValue = values[0]
		}
		select {
		case n := <- c1:
			values = append(values,n)
		case n := <- c2:
			values = append(values,n)
		case activeWork <- activeValue:
			values = values[1:]
		case <-time.After(800*time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("the queue len = ",len(values))
		case <-t:
			fmt.Println("bye")
			return
		}
	}
}
