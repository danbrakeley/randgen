package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

const (
	limit = 10 * 1024 * 1024 * 1024
)

func main() {
	f, err := os.Create("rando.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	n, err := io.Copy(f, io.LimitReader(rand.Reader, limit))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d random bytes\n", n)
}
