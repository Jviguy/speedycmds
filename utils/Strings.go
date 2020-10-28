package utils

import "strings"

func FindClosest(search string,array []string) string {
	var x []int
	x = make([]int , 0,len(array))
	for k,val := range array{
		a := strings.Split(search,"")
		b := strings.Split(val,"")
		c := 0
		for i := 0; i < len(b); i++ {
			if a[i] == b[i]{
				c++
			}
		}
		x[k] = c
	}
	min:=x[0]

	for i:=0; i<len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return array[x[min]]
}
