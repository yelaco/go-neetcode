package arrays_hashing

import "fmt"

// Use bucket sort
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int, len(nums))
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num] += 1
	}

	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	res := make([]int, 0, k)

	for i := len(freq) - 1; i > -1; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

func TestTopKFrequent() {
	numsCases := [][]int{
		{1, 2, 2, 3, 3, 3},
		{7, 7},
		{5},
	}
	ks := []int{
		2,
		1,
		1,
	}

	fmt.Println("* Two Integer Sum")
	for i, nums := range numsCases {
		fmt.Printf("Input: nums = %v, k = %d\nOutput: %v\n", nums, ks[i], topKFrequent(nums, ks[i]))
		fmt.Println()
	}
}
