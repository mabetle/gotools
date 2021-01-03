package main

import "flag"

// test command run times.

var (
	n int
)

func main() {
	flag.IntVar(&n, "n", 10, "run times")
	flag.Parse()

}
