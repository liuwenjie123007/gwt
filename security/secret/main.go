package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

func normal() {
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.'")
	fmt.Printf("% x\n", h.Sum(nil))

	h = sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.'")
	fmt.Printf("% x\n", h.Sum(nil))

	h = md5.New()
	io.WriteString(h, "需要加密的密码")
	fmt.Printf("% x\n", h.Sum(nil))
}

func advance() {
	h := md5.New()
	io.WriteString(h, "需要加密的密码")

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	salt1 := "@#$%"
	salt2 := "^&*()"

	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
}

func professional() {
	// scrypt方案
	salt := "@#$%"
	dk, _ := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32)
	fmt.Println(dk)
}

func main() {
	normal()
	advance()
	professional()
}
