package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/waltonseymour/sieve"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please enter a max value for the sieve")
		os.Exit(-1)
	}

	max := os.Args[1]

	parsedMax, err := strconv.Atoi(max)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, sieve.New(parsedMax))
}
