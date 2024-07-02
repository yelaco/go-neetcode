package stack

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	pairs := make([][2]int, 0, len(position))

	for i := 0; i < len(position); i++ {
		pairs = append(pairs, [2]int{position[i], speed[i]})
	}

	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	fleetStack := []float64{}

	for _, pair := range pairs {
		fleetStack = append(fleetStack, float64(target-pair[0])/float64(pair[1]))
		n := len(fleetStack)

		/*
			If the number of steps needed is smaller
			=> gonna reach the ahead car and form a car fleet
			If the number of steps needed is bigger
			=> never reach the ahead fleet
			=> gonna form a new fleet
			===> The stack just gonna keep the head of the car fleets
		*/
		if n >= 2 && fleetStack[n-1] <= fleetStack[n-2] {
			fleetStack = fleetStack[:n-1]
		}
	}

	return len(fleetStack)
}

func TestCarFleet() {
	targets := []int{
		10,
		10,
	}
	positions := [][]int{
		{1, 4},
		{4, 1, 0, 7},
	}
	speeds := [][]int{
		{3, 2},
		{2, 2, 1, 1},
	}

	for i := 0; i < len(positions); i++ {
		fmt.Printf("Input: target = %v, position = %v, speed = %v\nOutput: %v\n", targets[i], positions[i], speeds[i], carFleet(targets[i], positions[i], speeds[i]))
		fmt.Println()
	}
}
