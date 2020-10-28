package utils

import (
	"github.com/agnivade/levenshtein"
)

func FindClosest(search string,array []string) string {
	var x []int
	x = make([]int , 0,len(array))
	for _,val := range array {
		c := levenshtein.ComputeDistance(search,val)
		x = append(x,c)
	}
	max := x[0]
	for i:=0; i<len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	return array[x[max]]
}
