package collection

import "fmt"

func RunArray() {
	// init
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a=", a)

	b := [...]int{1, 2, 3}
	fmt.Println("b=", b)

	// set
	a[1] = 100
	fmt.Println("a=", a)

	// get
	fmt.Println("get", a[0])

	// 2d
	var aaa [2][3][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				aaa[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("aaa=", aaa)

}
