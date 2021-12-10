package oo

import "fmt"

func RunInterface() {

	var x A = &X{100}
	var y A = Y(101)

	call_a(x)
	call_a(y)

	fmt.Println("----")

	hi := Hi{"hi~"}
	hello := Hello{"hello~"}
	doGreet(hi)
	doGreet(hello)
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

type Greet interface {
	greet()
}

type Hi struct {
	v string
}

type Hello struct {
	m string
}

func (hi Hi) greet() {
	fmt.Println("hi")
}

func (hello Hello) greet() {
	fmt.Println("hello")
}

func doGreet(g Greet) {
	g.greet()

	if hi, ok := g.(Hi); ok {
		fmt.Println("if hi v=", hi.v)
	}
	if hello, ok := g.(Hello); ok {
		fmt.Println("if hi v=", hello.m)
	}

	switch g.(type) {
	case Hi:
		fmt.Println("is hi")
	case Hello:
		fmt.Println("is hello")
	default:
		fmt.Println("default")
	}
}
