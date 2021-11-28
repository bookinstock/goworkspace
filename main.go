package main

import (
	"fmt"
	"goworkspace/app/foo"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main run")

	foo.Run()
}
