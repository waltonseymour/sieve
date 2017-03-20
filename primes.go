package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func primes(maxInt int) {
	os.Create("out.csv")
	f, err := os.OpenFile("out.csv", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(f)
	var nums = make([]bool, maxInt, maxInt)
	for i := 2; i < maxInt; i++ {
		// false is unseen
		if nums[i] == false {
			fmt.Fprintln(writer, i)
			for j := i; j < maxInt; j += i {
				nums[j] = true
			}
		}
	}
	writer.Flush()
	f.Close()
}

func main() {
	args := os.Args
	if len(args) > 1 {
		i, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Parameter should be integer")
		} else {
			primes(i)
		}
	} else {
		primes(100)
	}
}
