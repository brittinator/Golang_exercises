package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTelephone(t *testing.T) {
	assert := assert.New(t)
	testTelly := telephone()
	testCase := testTelly(" test")
	assert.Equal(testCase, "I test")

}
