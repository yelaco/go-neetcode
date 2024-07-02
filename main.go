package main

import (
	arrays_hashing "github.com/minhquang053/neetcode/arrays-hashing"
	"github.com/minhquang053/neetcode/stack"
	two_pointers "github.com/minhquang053/neetcode/two-pointers"
)

func main() {
	// testArrayHashing()
	// testTwoPointers()
	testStack()
}

func testArrayHashing() {
	// arrays_hashing.TestDuplicate()
	// arrays_hashing.TestIsAnagram()
	// arrays_hashing.TestTwoSum()
	// arrays_hashing.TestAnagramGroups()
	// arrays_hashing.TestTopKFrequent()
	// arrays_hashing.TestEncodeDecodeStrings()
	// arrays_hashing.TestProductExceptSelf()
	arrays_hashing.TestLongestConsecutive()
	// arrays_hashing.TestValidSudoku()
}

func testTwoPointers() {
	// two_pointers.TestValidPalindrome()
	// two_pointers.TestTwoSumII()
	// two_pointers.TestThreeSum()
	// two_pointers.TestMaxArea()
	two_pointers.TestTrappingRainWater()
}

func testStack() {
	// stack.TestValidateParentheses()
	// stack.TestMinStack()
	// stack.TestEvalRPN()
	// stack.TestGenerateParentheses()
	// stack.TestDailyTemperatures()
	// stack.TestCarFleet()
	stack.TestLargestRectArea()
}
