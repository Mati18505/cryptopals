package main

func FixedXOR(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("length of a must be equal length of b")
	}

	result := make([]byte, len(a))

	for i, a := range a {
		result[i] = a ^ b[i]
	}

	return result
}