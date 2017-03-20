package main

import (
	"bufio"
	"fmt"
	"os"
)

func primes() {
	const maxInt uint = 1000
	var nums [maxInt]bool
	for i := uint(2); i < maxInt; i++ {
		// false is unseen
		if nums[i] == false {
			writeToFile(i)
			for j := i; j < maxInt; j += i {
				nums[j] = true
			}
		}
	}
}

func writeToFile(num uint) {
	f, err := os.OpenFile("primes.csv", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(f)
	fmt.Fprintln(writer, num)
	writer.Flush()
	defer f.Close()
}

func main() {
	primes()
}
