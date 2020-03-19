package main

import (
	"fmt"
	"time"
)

func main() {

	for i :=0; i<10; i++{
		go func(i int) {
			fmt.Printf("hello from gorutine %d\n",i)
		}(i)
	}
	time.Sleep(time.Microsecond)
}
