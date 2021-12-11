package dt

import (
	"fmt"
	"time"
)

// unix epoch
// sec 10
// milli 13
// micro 16
// nana 19

func RunEpoch() {
	var p = fmt.Println

	t0 := time.Now()
	p("t0=", t0)

	sec := t0.Unix()
	nano := t0.UnixNano()
	milli := nano / 1_000_000
	micro := nano / 1_000
	p("sec=", sec)
	p("milli=", milli)
	p("micro=", micro)
	p("nano=", nano)

	t1 := time.Unix(sec, 0)
	t2 := time.Unix(0, nano)
	p("t1= ", t1)
	p("t2= ", t2)
	p("t1 == t2", t1 == t2)
	p("t0 == t2", t0 == t2)
}
