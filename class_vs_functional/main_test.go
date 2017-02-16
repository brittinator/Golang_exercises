package main

import (
	"testing"
)

func BenchmarkMethod(b *testing.B) {
	var a A = "test"
	var i int = 123

	for n := 0; n < b.N; n++ {
		a.Demo(i)
	}
}

func BenchmarkFunction(b *testing.B) {
	var a A = "test"
	var i int = 123

	for n := 0; n < b.N; n++ {
		Demo(a, i)
	}
}
