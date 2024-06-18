package arrays_hashing

import "fmt"

func hasDuplicate(nums []int) bool {
	solMap := map[int]bool{}

	for _, num := range nums {
		if _, ok := solMap[num]; ok {
			return true
		} else {
			solMap[num] = true
		}
	}

	return false
}

func TestDuplicate() {
	numsList := [][]int{
		{1, 2, 3, 3},
		{1, 2, 3, 4},
	}

	fmt.Println("* Contains Duplicate")
	for _, nums := range numsList {
		fmt.Printf("Input: nums = %v\nOutput: %v\n", nums, hasDuplicate(nums))
		fmt.Println()
	}
}
