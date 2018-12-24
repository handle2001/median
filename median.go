// Package median provides functions for determining median values of a set
package median

import (
	"math"
	"sort"
)

// GetMedian returns the median value of a set
func GetMedian(s []int) []int {
	var median []int
	sort.Ints(s)
	slen := len(s)
	if slen%2 != 0 {
		mIndex := slen / 2
		median = append(median, s[mIndex])
	} else if slen%2 == 0 {
		x := float64(slen / 2)
		ceiling := int(math.Ceil(x))
		floor := int(math.Floor(x))
		m := (s[ceiling] + s[floor]) / 2
		median = append(median, m)
	}

	return median
}
