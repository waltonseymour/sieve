package main

import (
	"io"
	"os"

	"github.com/waltonseymour/sieve"
)

func main() {
	io.Copy(os.Stdout, sieve.New(1000))
}
