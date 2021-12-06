package list

import "fmt"

func array() {
	fmt.Println("---array---")
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}

	fmt.Println("a=", a, len(a), a[0])
	fmt.Println("b=", b, len(b), b[0])

	for i, e := range a {
		fmt.Println("i=", i, "e=", e)
	}

	var arr2d [3][4]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			arr2d[i][j] = i * j
		}
	}
	fmt.Println("arr2d=", arr2d)
}
