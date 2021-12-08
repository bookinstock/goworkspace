package list

import "fmt"

func RunSlice() {
	// init make

	s := make([]int, 5)
	fmt.Println("s=", s)

	s = make([]int, 5, 10)
	fmt.Println("s=", s)

	// init
	s = []int{1, 2, 3, 4, 5}
	fmt.Println("s=", s)

	// set
	s[1] = 100
	fmt.Println("s=", s)

	// append
	s = append(s, 100)
	fmt.Println("s=", s)

	as := []int{-1, -2}
	s = append(s, as...)
	fmt.Println("s=", s)

	// slice it
	s1 := s[:]
	s2 := s[1:]
	s3 := s[:1]
	s4 := s[1:3]

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
	fmt.Println("s4=", s4)

	// 2d
	lvl1Len := 2
	lvl2Len := 3
	ss := make([][]int, lvl1Len)
	for i := 0; i < lvl1Len; i++ {
		ss[i] = make([]int, lvl2Len)
		for j := 0; j < lvl2Len; j++ {
			ss[i][j] = (i + 1) * (j + 1)
		}
	}
	fmt.Println("ss=", ss)

	// 3d
	lvl1Len = 2
	lvl2Len = 3
	lvl3Len := 4
	sss := make([][][]int, lvl1Len)
	for i := 0; i < lvl1Len; i++ {
		sss[i] = make([][]int, lvl2Len)
		for j := 0; j < lvl2Len; j++ {
			sss[i][j] = make([]int, lvl3Len)
			for k := 0; k < lvl3Len; k++ {
				sss[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("sss=", sss)

}
