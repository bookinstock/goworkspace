package concurrent

import "fmt"

func RunWorker() {
	workerPool()
}

// ! block chan need receiver first then sender
// ! range need close

func workerPool() {
	fmt.Println("==workerPool==")
	workerSize := 3
	taskSize := 10

	in := make(chan int)
	out := make(chan int)

	for i := 0; i < workerSize; i++ {
		go worker(i, in, out)
	}

	go func() {
		for i := 0; i < taskSize; i++ {
			in <- i
		}
		close(in)
	}()

	for i := 0; i < taskSize; i++ {
		r := <-out
		fmt.Println("result=", r)
	}
	close(out)
	fmt.Println("==workerPool==")
}

func worker(id int, in <-chan int, out chan<- int) {
	fmt.Println("worker start id=", id)
	for in_msg := range in {
		// in_msg := <-in
		out_msg := in_msg * 100
		fmt.Println("worker id=", id, "in_msg=", in_msg, "out_msg=", out_msg)
		out <- out_msg
	}
}
