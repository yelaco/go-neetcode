package stack

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	iStack := []int{}
	res := make([]int, n)

	for i := n - 1; i > -1; i-- {
		for len(iStack) > 0 && temperatures[i] >= temperatures[iStack[len(iStack)-1]] {
			iStack = iStack[:len(iStack)-1]
		}

		if len(iStack) > 0 {
			res[i] = iStack[len(iStack)-1] - i
		}
		iStack = append(iStack, i)
	}

	return res
}

func TestDailyTemperatures() {
	testCases := [][]int{
		{30, 38, 30, 36, 35, 40, 28},
		{22, 21, 20},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: temperatures = %v\nOutput: %v\n", testCase, dailyTemperatures(testCase))
		fmt.Println()
	}
}
