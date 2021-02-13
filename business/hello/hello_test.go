package hello_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tgmendes/go-service-template/business/hello"
	"testing"
)

func TestHello(t *testing.T) {
	greet := hello.Greet("Bob")

	assert.Equal(t, "Hello, Bob!", greet)
}
