package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func RunSync() {

	// justTest() // ???

	// multiWait()

	// atomicAdd()

	mutexCounter()
}

func multiWait() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("do n=", n)
		}(i)
	}

	wg.Wait()
}

func atomicAdd() {
	var n uint64

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddUint64(&n, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Println("n=", n)

}

type Counter struct {
	mu   sync.Mutex
	data map[string]int
}

func (c *Counter) incr(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[name] += 1
}

func mutexCounter() {
	var wg sync.WaitGroup
	c := &Counter{data: map[string]int{}}
	d := []string{"a", "b", "a", "a", "c", "b"}
	for _, e := range d {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			c.incr(name)
		}(e)
	}
	wg.Wait()
	fmt.Println("data=", c.data)
}

//// todo

func justTest() {
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
