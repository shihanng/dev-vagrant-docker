package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert.HTTPBodyContains(t, helloHandler, "GET", "/", nil, "Hello, World")
}
