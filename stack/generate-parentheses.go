package stack

import (
	"fmt"
	"strings"
)

func generateParentheses(n int) []string {
	sol := solution{
		stack: make([]string, 0, 2*n),
		res:   []string{},
		n:     n,
	}

	sol.backtrack(0, 0)

	return sol.res
}

type solution struct {
	stack []string
	res   []string
	n     int
}

func (s *solution) backtrack(openN, closedN int) {
	if openN == closedN && closedN == s.n {
		s.res = append(s.res, strings.Join(s.stack, ""))
		return
	}

	if openN < s.n {
		s.stack = append(s.stack, "(")
		s.backtrack(openN+1, closedN)
		s.stack = s.stack[:len(s.stack)-1]
	}

	if closedN < openN {
		s.stack = append(s.stack, ")")
		s.backtrack(openN, closedN+1)
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func TestGenerateParentheses() {
	testCases := []int{3, 1}

	for _, testCase := range testCases {
		fmt.Printf("Input: n = %v\nOutput: %v\n", testCase, generateParentheses(testCase))
		fmt.Println()
	}
}
