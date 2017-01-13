package main

import (
	"fmt"
	"time"
)

// spinner shows a wheel spinning indefinitely
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// fib is a recursive expensive operation
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	// this creates a go routine that calls spinner, doesn't wait
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // this will be slow
	fmt.Printf("\rFibonacci(%v) = %v\n", n, fibN)
}
