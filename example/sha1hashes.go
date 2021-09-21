package example

import (
	"crypto/sha1"
	"fmt"
)

func Sha1Hashes() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
