package oo

import "fmt"

func RunEmbedding() {
	var a interface{} = &E{
		Foo: Foo{1},
		Bar: Bar{2},
		v:   100,
	}

	fmt.Println("a=", a)

	if e, ok := a.(fb); ok {
		e.foo()
		e.bar()
		fmt.Println("---")
		gofb(e)
	}
}

type fb interface {
	foo()
	bar()
}

func gofb(e fb) {
	e.foo()
	e.bar()
}

type Foo struct {
	f int
}

func (f *Foo) foo() {
	fmt.Println("foo v=", f.f)
}

type Bar struct {
	b int
}

func (b *Bar) bar() {
	fmt.Println("bar v=", b.b)
}

type E struct {
	Foo
	Bar
	v int
}
