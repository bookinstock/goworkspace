package error

import (
	"errors"
	"fmt"
)

func Run() {
	fmt.Printf("e1=%v e1=%s t1=%T \n", errFoo, errFoo, errFoo)
	fmt.Printf("e2=%v e2=%s t2=%T \n", errBar, errBar, errBar)
	fmt.Printf("e3=%v e3=%s t4=%T \n", errC, errC, errC)

	var i interface{} = errC
	if err, ok := i.(*cError); ok {
		fmt.Println("err=", err, "ok=", ok)
	} else {
		fmt.Println("else")
	}
}

var errFoo = errors.New("err foo")
var errBar = errors.New("err bar")
var errC = &cError{a: 1, b: "s"}

type cError struct {
	a int
	b string
}

func (c *cError) Error() string {
	return fmt.Sprintf("c err %d %s", c.a, c.b)
}
