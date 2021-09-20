package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	Table.Init()
}

func TestConversion(t *testing.T) {

	message := "hello world"

	encoded := encodeMessage(message)
	decoded := decodeMessage(encoded)

	assert.Equal(t, message, decoded)
}
