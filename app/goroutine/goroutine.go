package goroutine

import (
	"fmt"
	"time"
)

func Run() {

	// go fnGo(1)

	// fnCall()

	// channelGo()

	// channelBuffer()

	// channelDirection()

	// singleWait()

	// multiWait()

	// channelSelect()

	// channelSelectTimeout()

	// selectDefault()

	// channelClose()

	// iterateChannel()

	// timer()

	ticker()

	time.Sleep(time.Second)

}

// go
func fnGo(a int) {
	fmt.Println("fnGo a=", a)
}

func fnCall() {
	go func(a int) {
		fmt.Println("fnCall a=", a)
	}(2)
}

// channel
func channelGo() {
	c := make(chan int, 1)
	c <- 3
	fmt.Println("channelGo c=", <-c)
}

func channelBuffer() {
	c := make(chan int)
	go func() {
		c <- 4
	}()
	fmt.Println("channelBuffer c=", <-c)
}

func channelDirection() {
	c := make(chan int)
	v := 100
	go channelIn(c, v)
	r := channelOut(c)
	fmt.Println("channelDirection v=", v, "r=", r)
}

func channelIn(c chan<- int, v int) {
	c <- v + 1
}

func channelOut(c <-chan int) int {
	return <-c
}

// wait
func singleWait() {
	c := make(chan int)

	go func(x int) {
		time.Sleep(time.Second)
		c <- x
	}(5)

	fmt.Println("singleWait r=", <-c)
}

func multiWait() {

}

// select
func channelSelect() {
	fmt.Println("channelSelect==")
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()

	go func() {
		c2 <- 2
		time.Sleep(time.Second * 2)
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("m1=", m1)
		case m2 := <-c2:
			fmt.Println("m2=", m2)
		}
	}
	fmt.Println("channelSelect==")
}

// select timeout
func channelSelectTimeout() {
	fmt.Println("channelSelectTimeout==")

	// test1
	c := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c <- 1
	}()
	select {
	case msg := <-c:
		fmt.Println("test1 msg =", msg)
	case <-time.After(time.Second):
		fmt.Println("test1 timeout")
	}

	// test2
	c = make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c <- 1
	}()
	select {
	case msg := <-c:
		fmt.Println("test2 msg =", msg)
	case <-time.After(time.Second * 2):
		fmt.Println("test2 timeout")
	}

	fmt.Println("channelSelectTimeout==")
}

// select non-blocking default
func selectDefault() {
	c1 := make(chan int)
	c2 := make(chan int)

	// recevie
	select {
	case <-c1:
		fmt.Println("test1 c1")
	default:
		fmt.Println("test1 default")
	}

	// send
	select {
	case c1 <- 100:
		fmt.Println("test2 c1")
	default:
		fmt.Println("test2 default")
	}

	// multi receive
	select {
	case <-c1:
		fmt.Println("test3 c1")
	case <-c2:
		fmt.Println("test3 c2")
	default:
		fmt.Println("test3 default")
	}

}

func channelClose() {
	fmt.Println("--channelClose--")
	job := make(chan int)
	done := make(chan bool)

	go func() {
		count := 0
		for {
			msg, more := <-job
			if more {
				fmt.Println("receive msg=", msg)
			} else {
				fmt.Println("done more=", more, "count=", count)
				done <- true
				return
			}
			count += 1
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send msg=", i)
		job <- i
	}
	close(job)

	<-done

	fmt.Println("--channelClose--")
}

func iterateChannel() {
	fmt.Println("---iterateChannel")

	// test1
	c := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c <- i
	}
	close(c)
	for e := range c {
		fmt.Println("c msg=", e)
	}

	// test2
	c2 := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			c2 <- i
		}
		close(c2)
	}()
	for e := range c2 {
		fmt.Println("c2 msg=", e)
	}

	fmt.Println("---iterateChannel")
}

func timer() {
	t := time.NewTimer(time.Second)
	a := <-t.C
	fmt.Println("a=", a)

	t1 := time.NewTimer(time.Second)
	t2 := time.NewTimer(time.Second * 2)

	go func() {
		r := <-t1.C
		fmt.Println("t1 r=", r)
		s1 := t1.Stop()
		s2 := t2.Stop()
		fmt.Println("t1, stop s1=", s1, "s2=", s2)
	}()

	go func() {
		r := <-t2.C
		fmt.Println("t2 r=", r)
		s1 := t2.Stop()
		s2 := t2.Stop()
		fmt.Println("t2, stop s1=", s1, "s2=", s2)
	}()

	time.Sleep(time.Second * 3)
}

func ticker() {
	fmt.Println("ticker")
	t := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case msg := <-t.C:
				fmt.Println("t msg=", msg)
			case <-done:
				t.Stop()
				fmt.Println("done")
				return
			}
		}
	}()

	time.Sleep(time.Second * 3)
	done <- true
	fmt.Println("ticker")
}
