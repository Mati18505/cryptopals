package main

import "encoding/hex"

func hexToByteSlice(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}