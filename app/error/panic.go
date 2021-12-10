package error

import (
	"fmt"
	"os"
)

func RunPanic() {

	// doPanic()

	// deferPanic()

	// deferExit()

	panicRecover()

}

// simple
func doPanic() {
	panic("WRONG!")
}

// run defer
func deferPanic() {
	defer fmt.Println("deferPanic")
	fmt.Println(1)
	panic("foo")
}

// no defer
func deferExit() {
	defer fmt.Println("deferExit")
	fmt.Println(2)
	os.Exit(1)
}

// recover
func panicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic =", r)
		}
	}()
	doPanic()
}
