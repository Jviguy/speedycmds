package utils

import (
	"github.com/agnivade/levenshtein"
	"sort"
)

func FindClosest(search string, array []string) string {
    list := []int{}
    for _, name := range array {
        list = append(list, levenshtein.ComputeDistance(search, name))
    }
    sort.Ints(list)
    return array[list[0]]
}
