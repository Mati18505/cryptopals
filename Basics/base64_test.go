package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexToBase64(t *testing.T) {
	hexTests := map[string]string {
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d":
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		//"abcd":
		//"q80=",
	}

	stringTests := map[string]string {
		//"Time is a created thing To say I dont have time is like saying I dont want to Lao Tzu":
		//"VGltZSBpcyBhIGNyZWF0ZWQgdGhpbmcgVG8gc2F5IEkgZG9udCBoYXZlIHRpbWUgaXMgbGlrZSBzYXlpbmcgSSBkb250IHdhbnQgdG8gTGFvIFR6dQ==",
	}


	for input, expected := range hexTests {
		inputByte := hexToByteSlice(input)
		require.Equal(t, string(EncodeBase64(inputByte)), expected)
	}
	for input, expected := range stringTests {
		inputByte := input
		require.Equal(t, string(EncodeBase64([]byte(inputByte))), expected)
	}
}