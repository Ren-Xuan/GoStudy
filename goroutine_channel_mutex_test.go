package main

import (
	"fmt"
	"time"
)

func goroutine_test() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}
	timeout:=time.After(2*time.Second)
	for i := 0; i < 5; i++ {
		select{
			case gopherID :=<-c：
			fmt.println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.println("ran out")
			return
		}

	}

}

func sleepGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000))*time.Millisecond)
	fmt.Println("...snore...")
	c <- id
}
