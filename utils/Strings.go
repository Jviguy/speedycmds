package utils

import (
	"strings"
)

func FindClosest(search string,array []string) string {
	var x []int
	x = make([]int , 0,len(array))
	for _,val := range array {
		a := strings.Split(search,"")
		b := strings.Split(val,"")
		c := 0
		for i := 0; i < len(a); i++ {
			if a[i] == b[i]{
				c++
			}
		}
		x = append(x,c)
	}
	min:=x[0]

	for i:=0; i<len(x); i++ {
		if x[i] > min {
			min = x[i]
		}
	}
	return array[x[min]]
}
