package arrays_hashing

import "fmt"

func anagramGroups(strs []string) (res [][]string) {
	groups := map[string][]string{}

	for _, str := range strs {
		count := [26]int{}
		for _, c := range str {
			count[c-'a'] += 1
		}

		// since the count for each char is the difference between anagrams
		// we can unique keys for the hashmap
		keyString := ""
		for _, v := range count {
			keyString += "#" + fmt.Sprint(v)
		}

		groups[keyString] = append(groups[keyString], str)
	}

	// this is just to convert the groups we got to return format
	for _, group := range groups {
		res = append(res, group)
	}

	return
}

func TestAnagramGroups() {
	testCases := [][]string{
		{"act", "pots", "tops", "cat", "stop", "hat"},
		{"x"},
		{""},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: strs = %v\nOutput: %v\n", testCase, anagramGroups(testCase))
		fmt.Println()
	}
}
