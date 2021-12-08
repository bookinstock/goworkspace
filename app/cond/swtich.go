package cond

import "fmt"

func RunSwitch() {
	switchCond()
	switchValue()
	switchInit()
	swtichType()
}

func switchCond() {
	fmt.Println("---switchCond---")
	a := 101
	switch {
	case a < 5:
		fmt.Println("< 5", "a=", a)
	case a < 10:
		fmt.Println("< 10", "a=", a)
	case a < 15, a > 100:
		fmt.Println("< 15 or > 100", "a=", a)
	default:
		fmt.Println("default", "a=", a)
	}
}

func switchValue() {
	fmt.Println("---switchValue---")
	a := 5
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

func switchInit() {
	fmt.Println("---switchInit---")
	switch a, b, c := 1, 2, 3; a + b + c {
	case 1, 2, 3:
		fmt.Println("1,2,3")
	case 4, 5, 6:
		fmt.Println("4,5,6")
	default:
		fmt.Println("default")
	}
}

func swtichType() {
	fmt.Println("---swtichType---")
	var t interface{} = 1
	switch t.(type) {
	case int:
		fmt.Println("type:int")
	case string:
		fmt.Println("type:string")
	default:
		fmt.Println("type:default")
	}
}
