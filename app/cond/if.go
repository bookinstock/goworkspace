package cond

import "fmt"

func RunIf() {
	a := 10

	if a < 5 {
		fmt.Println("if- 1 a=", a)
	} else if a < 10 {
		fmt.Println("if- 2 a=", a)
	} else {
		fmt.Println("if- 3 a=", a)
	}

	if x := 5; x == 5 {
		fmt.Println("x is 5 x = ", x)
	} else {
		fmt.Println("x else x=", x)
	}
}
