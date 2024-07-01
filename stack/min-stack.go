package stack

import "fmt"

type MinStack struct {
	stack    []int // keep track of values in the stack
	minTrack []int // keep track of the min value for each push operation
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if len(ms.minTrack) == 0 || val < ms.minTrack[len(ms.minTrack)-1] {
		ms.minTrack = append(ms.minTrack, val)
	} else {
		ms.minTrack = append(ms.minTrack, ms.minTrack[len(ms.minTrack)-1])
	}
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]

	// if the stack popped, the min value for the
	// previous version of the stack is also popped
	ms.minTrack = ms.minTrack[:len(ms.minTrack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minTrack[len(ms.minTrack)-1]
}

func minStackOperations(ops []string, nums [][]int) {
	var msk MinStack
	fmt.Print("[")
	for i, op := range ops {
		switch op {
		case "MinStack":
			msk = MinStack{}
			fmt.Print("null")
		case "push":
			msk.Push(nums[i][0])
			fmt.Print(",null")
		case "pop":
			msk.Pop()
		case "top":
			fmt.Printf(",%d", msk.Top())
		case "getMin":
			fmt.Printf(",%d", msk.GetMin())
		}
	}
	fmt.Println("]")
}

func TestMinStack() {
	testCases := [][]string{
		{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
	}
	nums := [][][]int{
		{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
	}

	for i, testCase := range testCases {
		fmt.Printf("Input: %v\nOutput:", testCase)
		minStackOperations(testCase, nums[i])
		fmt.Println()
	}
}
