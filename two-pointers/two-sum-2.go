package two_pointers

import "fmt"

func twoSumII(nums []int, target int) [2]int {
	i := 0
	j := len(nums) - 1

	for i < j {
		if sum := nums[i] + nums[j]; sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			return [2]int{i + 1, j + 1}
		}
	}

	return [2]int{}
}

func TestTwoSumII() {
	numsCases := [][]int{
		{1, 2, 3, 4},
		{2, 7, 11, 15},
		{-1, 0},
		{2, 3, 4},
	}
	targets := []int{
		3,
		9,
		-1,
		6,
	}

	fmt.Println("* Two Integer Sum II")
	for i, nums := range numsCases {
		fmt.Printf("Input: nums = %v, target = %d\nOutput: %v\n", nums, targets[i], twoSumII(nums, targets[i]))
		fmt.Println()
	}
}
