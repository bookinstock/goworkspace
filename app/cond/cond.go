package cond

import "fmt"

func init() {
	fmt.Println("cond init")
}

func Run() {
	fmt.Println("cond run")

	ifelse()
	ifinit()
	switchcond()
	swtichcase()
	switchcaseinit()
	switchcasetype()
}

func ifelse() {
	fmt.Println("---ifelse---")

	a := 2
	var s string
	if a == 1 {
		s = "a=1"
	} else if a == 2 {
		s = "a=2"
	} else {
		s = "default"
	}
	fmt.Println(s)
}

func ifinit() {
	fmt.Println("---ifinit---")
	if a := 1; a == 1 {
		fmt.Println("a=1")
	}
}

func switchcond() {
	fmt.Println("---switchcond---")
	a := 1
	switch {
	case a == 1:
		fmt.Println("a=1")
	case a == 2:
		fmt.Println("a=2")
	case a >= 3 && a <= 5:
		fmt.Println("a=3,4,5")
	default:
		fmt.Println("default")
	}
}

func swtichcase() {
	fmt.Println("---switchcase---")
	a := 2
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	case 3, 4, 5:
		fmt.Println("a=3,4,5")
	default:
		fmt.Println("default")
	}
}

func switchcaseinit() {
	fmt.Println("---switchcaseinit---")
	switch a := 1; a {
	case 1:
		fmt.Println("a=1")
	}
}

func switchcasetype() {
	fmt.Println("---switchcasetype---")
	var t interface{} = 1

	showType(t)
}

func showType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}
