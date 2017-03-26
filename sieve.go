package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func sieve(maxInt int, outFile string) {
	os.Create(outFile)
	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_APPEND, 0660)
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
	var maxInt int
	var outFile string

	app := cli.NewApp()
	app.Name = "Sieve"
	app.Version = "0.1.0"
	app.Usage = "cli tool for fast prime number generation"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate a list of all prime numbers < maxInt to outFile",
			Action: func(c *cli.Context) error {
				sieve(maxInt, outFile)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "out",
					Value:       "out.csv",
					Usage:       "Output file for results",
					Destination: &outFile,
				},
				cli.IntFlag{
					Name:        "maxInt",
					Value:       1000,
					Usage:       "Maximum number to generate primes up to",
					Destination: &maxInt,
				},
			},
		},
	}

	app.Run(os.Args)
}
