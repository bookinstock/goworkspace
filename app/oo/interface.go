package oo

import "fmt"

func RunInterface() {

	var x A = &X{100}
	var y A = Y(101)

	call_a(x)
	call_a(y)
}

type A interface {
	a()
}

type X struct {
	val int
}

type Y int

func (x *X) a() {
	fmt.Println("x a val=", x.val)
}

func (y Y) a() {
	fmt.Println("y a =", y)
}

func call_a(a A) {
	a.a()
}
