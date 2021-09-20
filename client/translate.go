package client

import (
	"regexp"
	"strings"
)

func addPunctuation(encodedMessage, lang string) string {
	splittedMessage := strings.Split(encodedMessage, string(sep))
	index := 0

	for index < len(splittedMessage) {
		r := splittedMessage[index][0]
		delta := int((r % 60) + 1)
		index += delta
		if index >= len(splittedMessage) {
			break
		}

		splittedMessage[index] += lang
		mod := index % 32
		switch mod {
		case 0, 1, 2, 3:
			splittedMessage[index] += "，"
		case 7:
			splittedMessage[index] += "。"
		case 8:
			splittedMessage[index] += "？"
		case 9:
			splittedMessage[index] += "！"
		case 10:
			splittedMessage[index] += "~"
		}
	}

	return lang + strings.Join(splittedMessage, string(sep)) + lang
}

func removePunctuation(s string) string {
	re, _ := regexp.Compile("[^\u200b\u200c\u200d\u200e\u200f]")
	return re.ReplaceAllString(s, "")
}
