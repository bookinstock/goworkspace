package str

import (
	"crypto/sha1"
	"fmt"
)

func RunSha1() {
	s := "sha1 string"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println("s=", s)
	// fmt.Println("bs=", bs)
	fmt.Printf("bs= %x", bs)
}
