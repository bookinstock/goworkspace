package main

import (
	"fmt"
	"goworkspace/app/cond"
	"goworkspace/app/foo"
	"goworkspace/app/list"
	"goworkspace/app/loop"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main run")

	foo.Run()

	loop.Run()
	cond.Run()

	list.Run()
}
