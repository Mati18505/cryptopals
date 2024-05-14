package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXORKeyFinder(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"

	inputByte := hexToByteSlice(input)
	key := FindSingleCharKeyXOR(inputByte)
	decoded := string(decodeSingleByteXOR(inputByte, key))

	t.Logf("%v: %s", key, decoded)
	require.Equal(t, int(key), 88)
	require.Equal(t, decoded, expected)
}