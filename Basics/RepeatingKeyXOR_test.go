package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptRepeatingKeyXOR(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	expectedS := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	key := "ICE"
	expected := hexToByteSlice(expectedS)

	result := EncryptRepeatingKeyXOR([]byte(input), []byte(key))
	require.Equal(t, string(expected), string(result))
}

func TestMultipleEncryptRepeatingKeyXOR(t * testing.T) {
	inputs := []string {
		"mstulczewski@gmail.com",
		"Mateusz4321",
	}
	key := "ICE"
	for _, input := range inputs {
		t.Log(string(EncryptRepeatingKeyXOR([]byte(input), []byte(key))))
	}
}