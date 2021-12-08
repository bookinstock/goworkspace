package fn

import "fmt"

func RunRecursion() {
	testFactor()
	testFib()
}

func testFactor() {
	for i := 1; i <= 5; i++ {
		fmt.Println("factor i=", i, "r=", factor(i))
	}
}

func testFib() {
	for i := 1; i <= 5; i++ {
		fmt.Println("factor i=", i, "r=", fib(i))
	}
}

func factor(n int) int {
	if n <= 1 {
		return n
	}

	return n * factor(n-1)
}

// 1 1 2 3 4
func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
