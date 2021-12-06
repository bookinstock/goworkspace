package list

import "fmt"

func slice() {
	fmt.Println("---slice--")

	s1 := make([]int, 2, 4)
	s1[0] = 1
	s1[1] = 2
	// s1[2] = 3 // error

	s2 := []int{1, 2, 3}

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	s3 := append(s1, 100)
	s3 = append(s3, s2...)
	fmt.Println("s3 = ", s3)

	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println(s4[:])
	fmt.Println(s4[0:5])
	fmt.Println(s4[:2])
	fmt.Println(s4[2:])
	fmt.Println(s4[2:3])

	var s3d [][][]int = make([][][]int, 2)

	for a := 0; a < 2; a++ {
		s3d[a] = make([][]int, 3)
		for b := 0; b < 3; b++ {
			s3d[a][b] = make([]int, 4)
			for c := 0; c < 4; c++ {
				s3d[a][b][c] = a * b * c
			}
		}
	}

	fmt.Println("s3d=", s3d)
}
