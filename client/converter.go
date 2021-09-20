package client

import (
	"encoding/base64"
	"strings"
)

func encodeMessage(message string) string {
	encodedString := base64.RawStdEncoding.EncodeToString([]byte(message))
	zeroLenChars := make([]string, len(encodedString))
	for i, c := range encodedString {
		index := strings.IndexRune(b64, c)
		char := Table.IndexOf(index)
		zeroLenChars[i] = char
	}
	return strings.Join(zeroLenChars, sep)
}

func decodeMessage(message string) string {
	zeroLenChars := strings.Split(message, sep)
	var retMessage string = ""
	for _, char := range zeroLenChars {
		index := Table.CharIndex(char)
		retMessage += string(b64[index])
	}
	bs, _ := base64.RawStdEncoding.DecodeString(retMessage)
	return string(bs)
}
