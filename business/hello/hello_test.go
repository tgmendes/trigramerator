package hello_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tgmendes/trigramerator/business/hello"
)

func TestHello(t *testing.T) {
	greet := hello.Greet("Bob")

	assert.Equal(t, "Hello, Bob!", greet)
}
