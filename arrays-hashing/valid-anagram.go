package arrays_hashing

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sWordCount := map[byte]int{}
	tWordCount := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sWordCount[s[i]] += 1
		tWordCount[t[i]] += 1
	}

	// Compare if two map have the same entries with same value
	for key, sval := range sWordCount {
		if tval, ok := tWordCount[key]; !ok || tval != sval {
			return false
		}
	}
	return true
}

func TestIsAnagram() {
	testCases := [][]string{
		{"racecar", "carrace"},
		{"jar", "jam"},
	}

	fmt.Println("* Valid Anagram")
	for _, testCase := range testCases {
		fmt.Printf("Input: s = %s, t = %s\nOutput: %v\n", testCase[0], testCase[1], isAnagram(testCase[0], testCase[1]))
		fmt.Println()
	}
}
