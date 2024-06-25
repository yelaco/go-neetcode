package two_pointers

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	sort.Ints(nums)

	for i, num := range nums {
		// either it starts with positive num (not solvable)
		// or we have found all the possible triplets
		if num > 0 {
			break
		}

		// prevent duplicate triplet solutions
		if i > 0 && num == nums[i-1] {
			continue
		}

		/* We have two pointers to the second first and the last element
		we don't have first pointer to the first element because `num` starts from first element */
		left := i + 1
		right := len(nums) - 1

		for left < right {
			threeSum := num + nums[left] + nums[right]
			if threeSum < 0 {
				left += 1
			} else if threeSum > 0 {
				right -= 1
			} else {
				res = append(res, []int{num, nums[left], nums[right]})
				left += 1
				right -= 1

				/*
					Skip all the duplicate left pointer element
					to avoid duplicate triplets
					[!] By doing this, we effectively filter out
					duplicate right elements too since all the
					left elements that can make 0 sum with
					the right element are eliminated
				*/
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
			}
		}
	}

	return res
}

func TestThreeSum() {
	testCases := [][]int{
		{-1, 0, 1, 2, -1, 4},
		{0, 1, 1},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: nums = %v\nOutput: %v\n", testCase, threeSum(testCase))
		fmt.Println()
	}
}
