package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("viraj")
	assert.Equal(t, result, "Hello world : viraj")
}

