package fn

import "fmt"

func RunLambda() {
	testCounter()

	testNumProcessor()

}

func testCounter() {
	c1, c2 := counter(), counter()

	for i := 0; i < 3; i++ {
		fmt.Println("c1= ", c1())
	}

	for i := 0; i < 5; i++ {
		fmt.Println("c2= ", c2())
	}
}

func counter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func testNumProcessor() {
	p1 := func(n int) int {
		return n + 1
	}
	p2 := func(n int) int {
		return n * n
	}
	for i := 0; i < 3; i++ {
		fmt.Println("process i=", i, "r=", numProcessor(i, p1, p2))
	}

}

func numProcessor(n int, fns ...func(int) int) int {
	for _, fn := range fns {
		n = fn(n)
	}
	return n
}
