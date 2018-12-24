package median

import (
	"math"
	"sort"
)

// Write a function that takes a set of unsorted values and finds their median.

func getMedian(s []int) int {
	var median int
	// To find the median of a set takes  steps:
	// 1. Sort the set from smallest to largest
	sort.Ints(s)
	// 2. Find the length of the set
	slen := len(s)
	// 3. If the length of the set is odd, adding one to the length and then
	//		dividing that number by two gives the index of the median.
	if slen%2 != 0 {
		mIndex := slen / 2
		median = s[mIndex]

		// 4. If the length of the set is even, a few more steps are needed:
	} else if slen%2 == 0 {
		//	4a. Divide the length by two, this gives a float
		x := float64(slen / 2)
		//	4b. Round the float to the nearest integer up and down
		ceiling := int(math.Ceil(x))
		floor := int(math.Floor(x))
		//	4c. Average the values found at the two indexes found in 4b
		median = (s[ceiling] + s[floor]) / 2
	}

	return median
}
