package main

import (
	"fmt"
	"goworkspace/app/oo"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main run")

	// foo.Run()

	// variable.Run()

	// cond.Run()

	// cond.RunIf()
	// cond.RunSwitch()

	// loop.Run()

	// list.Run()

	// hashmap.Run()

	// loop.RunRange()

	// fn.Run()
	// fn.RunLambda()
	// fn.RunRecursion()

	oo.RunStruct()
	oo.RunInterface()

	// redisx.RunExample1()
	// redisx.Run()
}
