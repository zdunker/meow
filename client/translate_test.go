package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovePunctuation(t *testing.T) {
	s := "a\u200bb\u200cc\u200d"
	ss := removePunctuation(s)
	assert.Equal(t, "\u200b\u200c\u200d", ss)
}

func TestCircle(t *testing.T) {
	s := "hello world"
	encoded := encodeMessage(s)
	lang := "旺"
	dogLang := addPunctuation(encoded, lang)

	clean := removePunctuation(dogLang)
	decoded := decodeMessage(clean)
	assert.Equal(t, s, decoded)
}
