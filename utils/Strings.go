package utils

import (
	"github.com/agnivade/levenshtein"
)

func FindClosest(search string, array []string) string {
    list := []int{}
    for _, name := range array {
        list = append(list, levenshtein.ComputeDistance(search, name))
    }
    min := list[0]
    for _, num := range list {
        if num < min {
            min = num
        }
    }
    var index int
    for i, val := range list {
        if val == min {
            index = i
            break    
        }
    }
    return array[index]
}
