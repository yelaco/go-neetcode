package arrays_hashing

import "fmt"

func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	longest := 0

	for _, num := range nums {
		set[num] = struct{}{}
	}

	for num := range set {
		if _, ok := set[num-1]; !ok {
			for length := 1; ; length++ {
				if longest < length {
					longest = length
				}
				if _, ok := set[num+length]; !ok {
					break
				}
			}
		}
	}

	return longest
}

func TestLongestConsecutive() {
	testCases := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: nums = %v\nOutput: %v\n", testCase, longestConsecutive(testCase))
		fmt.Println()
	}
}
