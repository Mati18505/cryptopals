package main

func EncryptRepeatingKeyXOR(input, key []byte) []byte {
	idx := 0
	result := make([]byte, 0)

	for _, v := range input {
		result = append(result, v^key[idx])

		idx++
		idx %= len(key)
	}
	return result
}