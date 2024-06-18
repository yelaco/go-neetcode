package arrays_hashing

import "fmt"

func twoSum(nums []int, target int) [2]int {
	diffMap := map[int]int{}

	for i, num := range nums {
		diff := target - num
		if prevIdx, ok := diffMap[num]; ok {
			return [2]int{prevIdx, i}
		}
		diffMap[diff] = i
	}

	return [2]int{}
}

func TestTwoSum() {
	numsCases := [][]int{
		{3, 4, 5, 6},
		{4, 5, 6},
	}
	targets := []int{
		7,
		10,
	}

	fmt.Println("* Two Integer Sum")
	for i, nums := range numsCases {
		fmt.Printf("Input: nums = %v, target = %d\nOutput: %v\n", nums, targets[i], twoSum(nums, targets[i]))
		fmt.Println()
	}
}
