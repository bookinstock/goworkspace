package dt

import (
	"fmt"
	"time"
)

func RunParse() {
	var p = fmt.Println

	t0 := time.Now()
	p("t0=", t0)

	t1, e := time.Parse(
		time.RFC3339,
		"1991-10-25T01:02:03+08:00",
	)
	p("t1=", t1, "e=", e)

	p(t1.Format("3:04PM"))
	p(t1.Format("Mon Jan _2 15:04:0 2006"))
	p(t1.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p("t2=", t2, "e=", e)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d+08:00\n",
		t1.Year(), t1.Month(), t1.Day(),
		t1.Hour(), t1.Minute(), t1.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
