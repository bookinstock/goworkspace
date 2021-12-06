package loop

import "fmt"

func init() {
	fmt.Println("loop init")
}

func Run() {
	fmt.Println("loop run")

	for_loop()
	for_range_i()
	for_range_e()
	continue_break()
	for_while()
}

var a = []int{1, 2, 3, 4, 5}

func for_loop() {
	fmt.Println("---for_loop---")
	for i := 0; i < len(a); i++ {
		fmt.Println("i=", i, "e=", a[i])
	}
}

func for_range_i() {
	fmt.Println("---for_range_i---")
	for i := range a {
		fmt.Println("i=", i, "e=", a[i])
	}
}

func for_range_e() {
	fmt.Println("---for_range_e---")
	for i, e := range a {
		fmt.Println("i=", i, "e=", e)
	}
}

func continue_break() {
	fmt.Println("---continue_break---")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 5 {
			break
		}
		fmt.Println("i=", i)
	}
}

func for_while() {
	fmt.Println("---for_while---")
	i := 0
	for i < 5 {
		fmt.Println("i=", i)
		i++
	}
}
