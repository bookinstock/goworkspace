package concurrent

import (
	"fmt"
	"time"
)

func RunRateLimit() {
	// rangeMessage()
	rangeTicker()
}

func rangeMessage() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	t := time.NewTicker(time.Second)
	for n := range c {
		<-t.C
		fmt.Println("n=", n)
	}
}

func rangeTicker() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	t := time.NewTicker(time.Second)
	for e := range t.C {
		n, ok := <-c
		if !ok {
			fmt.Println("close")
			return
		}
		fmt.Println("n=", n, e)
	}
}
