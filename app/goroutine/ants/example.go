package ants

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

func Example1() {
	p("---Example1---")
	defer ants.Release()

	runtimes := 100

	// use common pool
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		f2()
		wg.Done()
	}
	for i := 0; i < runtimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	p("running goroutines =", ants.Running())
	p("finish all tasks")

}

func Example2() {
	p("---Example2---")
	defer ants.Release()

	runtimes := 100

	// use common pool
	var wg sync.WaitGroup

	// pool, _ = ants.NewPool(10)

	syncCalculateSum := func() {
		f2()
		wg.Done()
	}
	for i := 0; i < runtimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	p("running goroutines =", ants.Running())
	p("finish all tasks")
}

func Example3() {
	p("---Example3---")
	defer ants.Release()

	runtimes := 100
	var wg sync.WaitGroup

	// use the pool with a function
	// set 10 to cap of goroutine pool and 1 sec for expired duration
	pool, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		f1(i)
		wg.Done()
	})
	defer pool.Release()
	for i := 0; i < runtimes; i++ {
		wg.Add(1)
		_ = pool.Invoke(int32(i))
	}
	wg.Wait()
	p("running goroutines = ", pool.Running())
	p("finish all tasks, result is =", sum)
}

var sum int32
var p = fmt.Println

func f1(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	p("run with n= ", n)
}
func f2() {
	time.Sleep(time.Millisecond * 10)
	p("hello world")
}
