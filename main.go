package main

import (
	"fmt"
	"goworkspace/app/db/redisx"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main run")

	// foo.Run()

	// loop.Run()
	// cond.Run()

	// list.Run()

	// redisx.RunExample1()
	redisx.Run()
}
