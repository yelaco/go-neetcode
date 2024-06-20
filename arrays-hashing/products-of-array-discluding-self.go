package arrays_hashing

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// The left side product for the starting element
	// and right side product for the ending element should be 1
	res[0] = 1
	rightProduct := 1

	// first we keep track of the left product in the result
	for i := 1; i < n; i++ {
		// current left product = previous left element * its left product
		res[i] = res[i-1] * nums[i-1]
	}

	// then we multiply the left product with the corresponding right product
	for i := n - 2; i > -1; i-- {
		// current right product = previous right element * its right product
		rightProduct *= nums[i+1]

		// product except self = left product * right product
		res[i] *= rightProduct
	}

	return res
}

func TestProductExceptSelf() {
	inputs := [][]int{
		{1, 2, 4, 6},
		{-1, 0, 1, 2, 3},
	}

	for _, input := range inputs {
		fmt.Printf("Input: %v\nOutput: %v\n", input, productExceptSelf(input))
		fmt.Println()
	}
}
