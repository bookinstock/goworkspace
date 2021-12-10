package concurrent

import "fmt"

func RunRange() {
	loopFor()
	fmt.Println("----")
	loopForRange()
}

func loopFor() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("r=", <-c)
	}

}

func loopForRange() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for e := range c {
		fmt.Println("r=", e)
	}
}
