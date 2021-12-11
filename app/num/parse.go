package num

import (
	"fmt"
	"strconv"
)

func RunParse() {
	p := fmt.Println

	f, _ := strconv.ParseFloat("1.23", 64)
	i, _ := strconv.ParseInt("123", 0, 64)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	u, _ := strconv.ParseUint("789", 0, 64)
	k, _ := strconv.Atoi("123")
	_, e := strconv.Atoi("wtf")

	p("f=", f)
	p("i=", i)
	p("d=", d)
	p("u=", u)
	p("k=", k)
	p("e=", e)
}
