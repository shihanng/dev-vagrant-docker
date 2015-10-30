package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	a := "Hello"
	b := "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")

}
