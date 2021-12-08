package main

import (
	"fmt"
	"goworkspace/app/cond"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main run")

	// foo.Run()

	// cond.Run()

	// cond.RunIf()
	cond.RunSwitch()

	// loop.Run()

	// list.Run()

	// hashmap.Run()

	// loop.RunRange()

	// fn.Run()

	// redisx.RunExample1()
	// redisx.Run()
}
