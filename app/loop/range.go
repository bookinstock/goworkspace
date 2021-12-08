package loop

import "fmt"

func RunRange() {

	a := []int{1, 2, 3, 4, 5}
	b := map[int]int{1: 11, 2: 22, 3: 33}
	c := "abc"

	fmt.Println("loop slice")

	for i := range a {
		fmt.Println("slice1 i=", i, "v=", a[i])
	}

	for i, v := range a {
		fmt.Println("slice2 i=", i, "v=", v)
	}

	fmt.Println("loop map")

	for k := range b {
		fmt.Println("map1 k=", k, "v=", b[k])
	}

	for k, v := range b {
		fmt.Println("map2 k=", k, "v=", v)
	}

	fmt.Println("loop string")
	for i, v := range c {
		fmt.Println("string i=", i, "v=", v)
	}
}
