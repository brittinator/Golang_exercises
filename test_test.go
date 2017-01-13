package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {
	var a int
	a = "Birfday"
	assert.Equal(t, a, a, "expected equal")
}
