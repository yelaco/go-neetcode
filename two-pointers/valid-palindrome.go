package two_pointers

import "fmt"

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if !isAlnum(s[i]) {
			i++
			continue
		}

		if !isAlnum(s[j]) {
			j--
			continue
		}

		if !compareIgnoreCase(s[i], s[j]) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func compareIgnoreCase(a byte, b byte) bool {
	if a < 'A' || b < 'A' {
		return a == b
	}

	if a > b {
		return (a - b) == 32
	} else if a < b {
		return (b - a) == 32
	}

	return true
}

func TestValidPalindrome() {
	testCases := []string{
		"Was it a car or a cat i saw",
		"tab a cat",
		"0P",
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: s = %v\nOutput: %v\n", testCase, isPalindrome(testCase))
		fmt.Println()
	}
}
