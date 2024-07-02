package two_pointers

import (
	"fmt"

	"github.com/minhquang053/neetcode/util"
)

func trapRainWater(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	l := 0
	r := len(heights) - 1
	leftMax := heights[l]
	rightMax := heights[r]
	res := 0

	for l < r {
		/*
			The pointers will move toward the highest
			of all vertical lines
		*/
		if leftMax < rightMax {
			l += 1
			leftMax = util.Max(leftMax, heights[l])
			// add the vertical gap to the sum area
			res += leftMax - heights[l]
		} else {
			r -= 1
			rightMax = util.Max(rightMax, heights[r])
			// add the vertical gap to the sum area
			res += rightMax - heights[r]
		}
	}

	return res
}

func TestTrappingRainWater() {
	testCases := [][]int{
		{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
		{1, 1},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: heights = %v\nOutput: %v\n", testCase, trapRainWater(testCase))
		fmt.Println()
	}
}
