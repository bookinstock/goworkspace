package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var p = fmt.Println
var filename = "./tmp/dat"

func Run() {
	read()
	write()
}

func read() {
	p("--read file--")
	d, _ := os.ReadFile(filename)
	p("d=", string(d))

	p("-- open --")
	f, _ := os.Open(filename)
	b := make([]byte, 4)
	n, _ := f.Read(b)
	p("n=", n)
	p("b=", b)
	p("b=", string(b[:n]))

	o, _ := f.Seek(4, 0)
	b = make([]byte, 3)
	n, _ = f.Read(b)
	p("o=", o, "n=", n)
	p("s=", string(b[:n]))

	o, _ = f.Seek(4, 0)
	b = make([]byte, 3)
	n, _ = io.ReadAtLeast(f, b, 3)
	p("o=", o, "n=", n)
	p("s=", string(b[:n]))

	_, _ = f.Seek(0, 0)
	r := bufio.NewReader(f)
	b, _ = r.Peek(5)
	p("s=", string(b))

	f.Close()
}

func write() {
	d := []byte("foo\nbar\n")
	_ = os.WriteFile(filename, d, 0644)

	f, _ := os.Create("./tmp/new")
	defer f.Close()

	d = []byte{115, 111, 109, 101, 10}
	n, _ := f.Write(d)
	print("n=", n)

	f.Sync()

	w := bufio.NewWriter(f)
	n, _ = w.WriteString("buffered\n")
	p("n=", n)

	w.Flush()
}
