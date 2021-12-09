package main

import (
	"fmt"
	"goworkspace/app/goroutine"
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

	// oo.RunStruct()
	// oo.RunInterface()
	// oo.RunEmbedding()

	// error.Run()

	goroutine.Run()

	// redisx.RunExample1()
	// redisx.Run()
}
