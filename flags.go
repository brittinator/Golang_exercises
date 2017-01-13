package main

import (
	"flag"
	"fmt"
)

func main() {
	// Register Int flag.
	count := flag.Int("count", 5, "count of iterations")
	// Parse the flags.
	flag.Parse()

	// Print the argument.
	fmt.Println("Argument", *count)

	// Get int from the Int pointer.
	value := *count
	fmt.Println("Value", value)
}
