package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindXorString(t *testing.T) {
	file, err := os.ReadFile("xor.txt")
	expected := []byte("Now that the party is jumping\n")
	require.NoError(t, err)

	result := FindXORStringInStrings(file)

	require.Equal(t, result, expected)
}