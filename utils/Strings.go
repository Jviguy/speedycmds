package utils

import (
	"github.com/agnivade/levenshtein"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func FindClosest(target string, list []string) string {
	fuz := FindFuzzy(target, list)
	lev := FindLevenshtein(target, list)
	if  fuz == nil || *fuz == *lev {
		return *lev
	}
	return "`" + *lev + "`" + " or `" + *fuz + "`?"
}

func FindFuzzy(target string, list []string) *string {
	fuz := fuzzy.Find(target, list)
	if len(fuz) > 0 {
		return &fuz[0]
	}
	return nil
}

func FindLevenshtein(target string, list []string) *string {
	results := []int{}
	for _, name := range list {
		results = append(results, levenshtein.ComputeDistance(target, name))
	}
	min := results[0]
	for _, num := range results {
		if num < min {
			min = num
		}
	}
	var index int
	for i, val := range results {
		if val == min {
			index = i
			break
		}
	}
	return &list[index]
}
