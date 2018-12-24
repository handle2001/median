package main

import (
	"fmt"
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

func main() {
	set := []int{90, 92, 93, 88, 95, 88, 97, 87, 98}
	median := getMedian(set)
	fmt.Printf("The median is %d\n", median)
}

// TODO:
// 1. Setup RESTful API to accept HTTP request containing set of values in JSON
// 2. Format and return JSON object containing the median value of the set.
// 3. Build a container that runs the median app
// 4. Run web benchmarks against the container
// 5. Deploy the container to a Kubernetes cluster
