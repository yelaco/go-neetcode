package two_pointers

import (
	"fmt"

	"github.com/minhquang053/neetcode/utils"
)

func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	res := 0

	for l < r {
		// area = shorter vertical line * (right - left)
		res = utils.Max(utils.Min(heights[l], heights[r])*(r-l), res)

		/*
			By keeping the higher vertical line,
			we know that the area with width (r-l)
			will always be the largest possible for
			each width as the width gradually decrease
		*/
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
		/*
			The other containers with the same width
			form from one of the vertical lines
			which are left out will always have less
			area since the line is shorter than lines
			at left and right pointer positions.
		*/
	}

	return res
}

func TestMaxArea() {
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: heights = %v\nOutput: %v\n", testCase, maxArea(testCase))
		fmt.Println()
	}
}
