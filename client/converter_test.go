package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp() {
	MakeTable()
}

func TestConversion(t *testing.T) {
	MakeTable()

	message := "hello world"

	encoded := encodeMessage(message)
	decoded := decodeMessage(encoded)

	assert.Equal(t, message, decoded)
}
