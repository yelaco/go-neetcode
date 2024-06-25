package stack

import "fmt"

func isValid(s string) bool {
	bracketStack := make([]rune, 0, len(s))
	bracketMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, ch := range s {
		n := len(bracketStack)
		if n > 0 {
			if bracket, ok := bracketMap[bracketStack[n-1]]; ok {
				if bracket == ch {
					bracketStack = bracketStack[:n-1]
					continue
				}
			} else {
				return false
			}
		}
		bracketStack = append(bracketStack, ch)
	}

	return len(bracketStack) == 0
}

func TestValidateParentheses() {
	testCases := []string{
		"[]",
		"([{}])",
		"[(])",
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: s = %v\nOutput: %v\n", testCase, isValid(testCase))
		fmt.Println()
	}
}
