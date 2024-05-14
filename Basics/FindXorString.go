package main

import (
	"bytes"
)

// Finds one xored string in strings stored as hex divided by '\r\n' and returns it after decrypting.
func FindXORStringInStrings(input []byte) []byte {
	best := struct {
		score float32
		decoded []byte
	}{}

	lines := bytes.Split(input, []byte("\r\n"))

	for _, line := range lines {
		bytes := hexToByteSlice(string(line))
		key := FindSingleCharKeyXOR(bytes)
		decoded := decodeSingleByteXOR(bytes, key)
		score := englishPlaintextScore(string(decoded))

		if score > best.score {
			best.score = score
			best.decoded = decoded
		}
	}
	return best.decoded
}