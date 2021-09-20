package client

const (
	b64   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/="
	sep   = "\u200b"
	codes = "\u200c\u200d\u200e\u200f"
)

var (
	Table lookupTable
)

type lookupTable struct {
	codes []string
	m     map[string]int
}

func (t *lookupTable) init() {
	Table.codes = []string{}
	Table.m = make(map[string]int)
}

func (t *lookupTable) Length() int {
	return len(t.codes)
}

func (t *lookupTable) Contains(s string) bool {
	_, exists := t.m[s]
	return exists
}

func (t *lookupTable) Push(s string) {
	t.codes = append(t.codes, s)
	t.m[s] = len(t.codes) - 1
}

func (t *lookupTable) IndexOf(i int) string {
	return t.codes[i]
}

func (t *lookupTable) CharIndex(char string) int {
	return t.m[char]
}
