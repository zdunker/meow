package client

func MakeTable() {
	Table.init()
	b64Length := len(b64)

	for Table.Length() < b64Length {
		tmpLen := Table.Length()
		for _, _code := range codes {
			if !Table.Contains(string(_code)) {
				if Table.Length() >= b64Length {
					break
				}
				Table.Push(string(_code))
			}
			for i := 0; i < tmpLen; i++ {
				if Table.Length() >= b64Length {
					break
				}
				newCode := string(_code) + Table.IndexOf(i)
				if !Table.Contains(newCode) {
					Table.Push(newCode)
				}
			}
		}
	}
}
