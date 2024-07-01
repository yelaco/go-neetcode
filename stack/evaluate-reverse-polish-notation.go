package stack

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	tokStack := []int{}
	opMap := map[string]func(int, int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"*": func(a, b int) int {
			return a * b
		},
		"/": func(a, b int) int {
			return a / b
		},
	}

	for _, tok := range tokens {
		if f, ok := opMap[tok]; ok {
			// pop 2 last elements from stack
			a := tokStack[len(tokStack)-2]
			b := tokStack[len(tokStack)-1]

			// push the result into stack
			tokStack = tokStack[:len(tokStack)-1]
			tokStack[len(tokStack)-1] = f(a, b)
		} else {
			// if not operator, push the number into stack
			tokInt, _ := strconv.Atoi(tok)
			tokStack = append(tokStack, tokInt)
		}
	}

	return tokStack[0]
}

func TestEvalRPN() {
	testCases := [][]string{
		{"2", "1", "+", "3", "*"},
		{"4", "13", "5", "/", "+"},
		{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: tokens = %v\nOutput: %v\n", testCase, evalRPN(testCase))
		fmt.Println()
	}
}
