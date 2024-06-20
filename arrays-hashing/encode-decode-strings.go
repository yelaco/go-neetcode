package arrays_hashing

import (
	"fmt"
	"strconv"
)

func encode(words []string) string {
	encoded := ""
	for _, word := range words {
		encoded += fmt.Sprint(len(word)) + "#" + word
	}

	return encoded
}

func decode(s string) []string {
	decoded := []string{}
	i := 0

	for i < len(s) {
		// check the number for length until we reach a #
		// since the word with length > 10 will be of multiple digits
		j := i
		for s[j] != '#' {
			j += 1
		}
		wordLength, _ := strconv.Atoi(s[i:j])
		decoded = append(decoded, s[j+1:j+1+wordLength])
		i = j + 1 + wordLength
	}

	return decoded
}

func TestEncodeDecodeStrings() {
	inputs := [][]string{
		{"neet", "code", "love", "you", "this-is-a-long-word"},
		{"we", "say", ":", "yes"},
	}

	for _, input := range inputs {
		encoded := encode(input)
		decoded := decode(encoded)
		fmt.Printf("Input: %v\nEncoded: %v\nDecoded: %v\n", input, encoded, decoded)
		fmt.Println()
	}
}
