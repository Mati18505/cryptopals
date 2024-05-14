package main

import (
	"math"
	"regexp"
	"unicode"
)

func FindSingleCharKeyXOR(input []byte) byte {
	best := struct {
		key byte
		score float32
	}{}

	for i := range math.MaxUint8 {
		decoded := decodeSingleByteXOR(input, byte(i))
		score := englishPlaintextScore(string(decoded))

		if score > best.score {
			best.key = byte(i)
			best.score = score
		}
		
	}
	return best.key
}

var letterFrequencies = map[rune]float32{
    'A': 8.167,
    'B': 1.492,
    'C': 2.782,
    'D': 4.253,
    'E': 12.702,
    'F': 2.228,
    'G': 2.015,
    'H': 6.094,
    'I': 6.966,
    'J': 0.153,
    'K': 0.772,
    'L': 4.025,
    'M': 2.406,
    'N': 6.749,
    'O': 7.507,
    'P': 1.929,
    'Q': 0.095,
    'R': 5.987,
    'S': 6.327,
    'T': 9.056,
    'U': 2.758,
    'V': 0.978,
    'W': 2.360,
    'X': 0.150,
    'Y': 1.974,
    'Z': 0.074,
}

var text = regexp.MustCompile("^[a-zA-Z ]$")

func isAlphabetic(s string) bool {
	return text.MatchString(s)
}

func englishPlaintextScore(input string) float32 {
	var result float32 = 0 

	for _, r := range input {
		if isAlphabetic(string(r)) {
			result += letterFrequencies[unicode.ToUpper(r)]
			continue
		} else {
			result -= 10.0
		}
	}

	return result
}

func decodeSingleByteXOR(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i, v := range input {
		result[i] = v ^ key
	}
	return result
}