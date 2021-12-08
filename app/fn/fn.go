package fn

import "fmt"

func Run() {
	fmt.Println("f1=", f1(1, 2))

	fmt.Println("f2=", f2(1, 2, 3))

	r1, r2 := f3(1, 2)
	fmt.Println("f3=", r1, r2)

	r1, r2 = f4(1, 2)
	fmt.Println("f4=", r1, r2)
}

func f1(a int, b int) int {
	return a + b
}

func f2(a, b, c int) int {
	return a + b + c
}

func f3(a, b int) (int, int) {
	return a + b, a - b
}

func f4(a, b int) (r1, r2 int) {
	r1 = a + b
	r2 = a - b
	return
}
