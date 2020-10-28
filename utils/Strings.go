package utils

import (
	"github.com/agnivade/levenshtein"
	"sort"
)

func FindClosest(search string,array []string) string {
	var x []int
	x = make([]int, len(array))
	for k, val := range array {
		c := levenshtein.ComputeDistance(search, val)
		x[k] = c
	}
	y := x
	sort.Ints(y)
	k := sort.SearchInts(x,y[0])
	return array[k]
}
