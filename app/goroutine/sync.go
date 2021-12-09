package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func RunSync() {
	wg := new(sync.WaitGroup)
	messages := make(chan string)
	for x := 1; x <= 5; x++ {
		wg.Add(1)
		go sayhello(x, wg, &messages)
	}

	// go func() {
	fmt.Println("--")
	// for msg := range messages {
	// 	fmt.Println(msg)
	// }
	for x := 1; x <= 5; x++ {
		fmt.Println("xxx=", <-messages)
	}

	close(messages)
	// }()

	wg.Wait()

}

func sayhello(count int, wg *sync.WaitGroup, messages *chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(1000))
	*messages <- fmt.Sprintf("hello: %d", count)
}
