package client

const (
	b64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/="
	sep = '\u200b'
)

var (
	codes = [][]rune{{'\u200c'}, {'\u200d'}, {'\u200e'}, {'\u200f'}}
	Table lookupTable
)

type lookupTable struct {
	codes [][]rune
	m     map[string]int
}

func (t *lookupTable) Length() int {
	return len(t.codes)
}

func (t *lookupTable) Contains(rs []rune) bool {
	s := string(rs)
	_, exists := t.m[s]
	return exists
}

func (t *lookupTable) Push(rs []rune) {
	t.codes = append(t.codes, rs)
	s := string(rs)
	t.m[s] = len(t.codes) - 1
}

func (t *lookupTable) IndexOf(i int) []rune {
	return t.codes[i]
}

func (t *lookupTable) CharIndex(char []rune) int {
	s := string(char)
	return t.m[s]
}

func (t *lookupTable) Init() {
	Table.codes = make([][]rune, 0, len(b64))
	Table.m = make(map[string]int)

	b64Length := len(b64)

	for Table.Length() < b64Length {
		tmpLen := Table.Length()
		for _, _code := range codes {
			if !Table.Contains(_code) {
				if Table.Length() >= b64Length {
					break
				}
				Table.Push(_code)
			}
			for i := 0; i < tmpLen; i++ {
				if Table.Length() >= b64Length {
					break
				}
				newCode := append(_code, Table.IndexOf(i)...)
				if !Table.Contains(newCode) {
					Table.Push(newCode)
				}
			}
		}
	}
}
