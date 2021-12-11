package variable

import "fmt"

func Run() {
	a := 1
	aa := &a
	fmt.Printf("a=%v t=%T\n", a, a)
	fmt.Printf("aa=%v t=%T\n", aa, aa)

	f1(&a)
	fmt.Println("a=", a)

	a = *f2()
	fmt.Println("a=", a)
}

func f1(x *int) {
	*x += 1
}

func f2() *int {
	x := 1
	return &x
}
