package main

// TODO: Padding (= at the end).
func EncodeBase64(input []byte) []byte {
	for len(input)%3 != 0 {
		input = append(input, 0)
	}

	// N/3 Groups each 3 bytes.
	groups := make([][3]byte, len(input)/3)

	n := 0
	for i := range groups {
		groups[i] = [3]byte{input[n], input[n+1], input[n+2]}
		n += 3
	}

	// Each group gives 4 result bytes.
	result := make([]byte, len(groups)*4)

	for i, g := range groups {
		// WTF is that???
		var a [4]byte
		a[0] = (g[0] & 0b11111100) >> 2

		a[1] = (g[0] & 0b11) << 4
		a[1] |= (g[1] & 0b11110000) >> 4

		a[2] = (g[1] & 0b00001111) << 2
		a[2] |= (g[2] & 0b11000000) >> 6

		a[3] = g[2] & 0b00111111

		result[i*4] = a[0]
		result[i*4+1] = a[1]
		result[i*4+2] = a[2]
		result[i*4+3] = a[3]
	}

	for i, b := range result {
		if b < 26 {
			result[i] += 65
		} else if b < 52 {
			result[i] += 71
		} else if b < 62 {
			result[i] -= 4
		} else if b == 62 {
			result[i] = '+'
		} else if b == 63 {
			result[i] = '/'
		}
	}

	return result
}