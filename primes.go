package main

import (
	"bufio"
	"fmt"
	"os"
)

func primes() {
	os.Create("out.csv")
	f, err := os.OpenFile("out.csv", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(f)

	const maxInt uint = 10000000
	var nums [maxInt]bool
	for i := uint(2); i < maxInt; i++ {
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
	primes()
}
