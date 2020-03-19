package main

import (
"fmt"
	"sync"
)

func worker(c chan  int){
	for{
		n := <-c //收channel数字
		fmt.Println(n)
	}
}


type workers struct {
	in chan int
	done chan bool
}


func doWork(id int , c chan  int, done chan bool){
	for n :=range c {
		/*n, ok :=<-c
		if !ok {
			break
		}*/
		fmt.Printf("Worker %d Received %c\n",id,n)
		go func() {
			done<-true
		}()

	}
}


//返回一个channel
func CreateWorker(id int ) workers {
	w := workers{
		 in :make(chan int),
		 done :make(chan bool),
	}

	go doWork(id,w.in, w.done)
	return w
}


func chanDemo(){
	//var c  chan  int // c==nil
	var workerx  [10]workers

	for i:=0;i<10;i++{
		//go worker(c)
		workerx[i] =CreateWorker(i)
	}

	for i, worker :=range workerx{
		worker.in <-'a'+i

	}
	for i , worker :=range workerx{
		worker.in <-'A'+i

	}
	//c <-1 //发channel数字 chan <- int
	//wait for all of then
	for _, worker :=range workerx{
		<- worker.done
		<- worker.done
	}
}
/////////////////////////////////////////////////////


type workersx struct {
	in chan int
	wg *sync.WaitGroup

}

func doWorkx(id int , c chan  int, wg *sync.WaitGroup){
	for n :=range c {
		/*n, ok :=<-c
		if !ok {
			break
		}*/
		fmt.Printf("Worker %d Received %c\n",id,n)
		go func() {
			wg.Done()
		}()

	}
}
//返回一个channel
func CreateWorkerx(id int,wg *sync.WaitGroup ) workersx {
	w := workersx{
		in :make(chan int),
		wg :wg,
	}

	go doWorkx(id,w.in, w.wg)
	return w
}


func chanDemox(){
	//var c  chan  int // c==nil
	var wg sync.WaitGroup
	var workerx  [10]workersx

	for i:=0;i<10;i++{
		//go worker(c)
		workerx[i] =CreateWorkerx(i,&wg)
	}

	wg.Add(20)//20 tasks

	for i, worker :=range workerx{
		worker.in <-'a'+i

	}
	for i , worker :=range workerx{
		worker.in <-'A'+i

	}
	wg.Done()
	wg.Wait()
	//c <-1 //发channel数字 chan <- int
	}


func main() {
	fmt.Println("channel as first-class citieszon")
	//chanDemo()
	chanDemox()
	}
