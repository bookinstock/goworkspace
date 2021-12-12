package collection

import "fmt"

func RunIterator() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("s=", s)

	fmt.Println("map=", Map(s, func(e int) int { return e + 1 }))

	fmt.Println("filter=", Filter(s, func(e int) bool { return e%2 == 0 }))

	fmt.Println("index=", Index(s, 3), Index(s, 100))

	fmt.Println("include=", Include(s, 3), Include(s, 100))

	fmt.Println("any=", Any(s, func(e int) bool { return e == 1 }), Any(s, func(e int) bool { return e == 100 }))

	fmt.Println("all=", All(s, func(e int) bool { return e == 1 }), All(s, func(e int) bool { return e < 100 }))

}

// index
func Index(s []int, v int) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

// include
func Include(s []int, v int) bool {
	return Index(s, v) != -1
}

// any
func Any(s []int, fn func(int) bool) bool {
	for _, e := range s {
		if fn(e) {
			return true
		}
	}
	return false
}

// all
func All(s []int, fn func(int) bool) bool {
	for _, e := range s {
		if !fn(e) {
			return false
		}
	}
	return true
}

// filter
func Filter(s []int, fn func(int) bool) []int {
	r := []int{}
	for _, e := range s {
		if fn(e) {
			r = append(r, e)
		}
	}
	return r
}

// map
func Map(s []int, fn func(int) int) []int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = fn(e)
	}
	return r
}
