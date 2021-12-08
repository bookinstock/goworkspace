package hashmap

import "fmt"

func Run() {

	// init make
	m := make(map[string]int, 3)
	fmt.Println("m=", m)

	// init
	m = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("m=", m)

	// get
	v1 := m["a"]
	v2 := m["x"]
	fmt.Println("v1=", v1)
	fmt.Println("v2=", v2)
	v1, ex1 := m["a"]
	v2, ex2 := m["x"]
	fmt.Println("v1=", v1, "ex1=", ex1)
	fmt.Println("v2=", v2, "ex2=", ex2)

	// set
	m["a"] += 1
	m["x"] += 1
	m["y"] = 100
	fmt.Println("m= set", m)

	// delete
	delete(m, "y")
	delete(m, "z")
	fmt.Println("m= delete", m)

	// len
	fmt.Println("len=", len(m))

	// loop

	for k := range m {
		fmt.Println("k=", k, "v=", m[k])
	}

	fmt.Println("---")

	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
	}

}
