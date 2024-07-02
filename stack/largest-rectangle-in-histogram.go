package stack

import (
	"fmt"

	"github.com/minhquang053/neetcode/util"
)

/*
Monotoic increasing stack
==> Keep the stack increasing until meet a smaller number
==> pop until that number is bigger than the top of the stack
==> while popping, calculate the maxArea possible for the popped elements.
==> After that, with the monotoic stack, we calculate the possible area for elements
*/
func largestRectangleArea(heights []int) int {
	maxArea := 0
	pStack := [][2]int{} // pair (height, index)

	for i, h := range heights {
		start := i
		for len(pStack) > 0 && pStack[len(pStack)-1][0] > h {
			rect := pStack[len(pStack)-1]
			pStack = pStack[:len(pStack)-1]
			maxArea = util.Max(maxArea, rect[0]*(i-rect[1]))
			start = rect[1]
		}
		pStack = append(pStack, [2]int{h, start})
	}

	for _, pair := range pStack {
		maxArea = util.Max(maxArea, pair[0]*(len(heights)-pair[1]))
	}

	return maxArea
}

func TestLargestRectArea() {
	testCases := [][]int{
		{7, 1, 7, 2, 2, 4},
		{1, 3, 7},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: heights = %v\nOutput: %v\n", testCase, largestRectangleArea(testCase))
		fmt.Println()
	}
}
