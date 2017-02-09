package main

import (
	"testing"
)

func TestLoopThrough(t *testing.T) {

}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopThrough(1000)
	}
}
