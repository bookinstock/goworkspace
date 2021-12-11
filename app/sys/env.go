package sys

import (
	"fmt"
	"os"
	"strings"
)

var p = fmt.Println

func RunEnv() {
	os.Setenv("FOO", "1")

	p("foo:", os.Getenv("FOO"))
	p("bar:", os.Getenv("BAR"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		p(pair[0])
	}
}
