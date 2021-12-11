package num

import (
	"fmt"
	"math/rand"
	"time"
)

func RunRand() {
	var p = fmt.Println

	for i := 0; i < 5; i++ {
		p("rand int=", rand.Intn(100))
	}

	for i := 0; i < 5; i++ {
		p("rand float=", rand.Float32())
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < 5; i++ {
		p("r int=", r.Intn(10))
	}

	r1 := rand.New(rand.NewSource(42))
	r2 := rand.New(rand.NewSource(42))
	for i := 0; i < 5; i++ {
		p("r1 int=", r1.Intn(10))
	}
	for i := 0; i < 5; i++ {
		p("r2 int=", r2.Intn(10))
	}

}
