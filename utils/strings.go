package utils

import (
	"github.com/agnivade/levenshtein"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

// FindClosest finds the closest string in the list provided to target and returns it.
func FindClosest(target string, list []string) string {
	fuz, ok := FindFuzzy(target, list)
	lev := FindLevenshtein(target, list)
	if !ok {
		return lev
	}
	return "`" + lev + "`" + " or `" + fuz + "`?"
}

// FindFuzzy finds a target in a list of strings using the Fuzzy matching algorithm.
func FindFuzzy(target string, list []string) (string, bool) {
	fuz := fuzzy.Find(target, list)
	if len(fuz) > 0 {
		return fuz[0], true
	}
	return "", false
}

// FindLevenshtein finds a target in a list of strings using levenshtein distance.
func FindLevenshtein(target string, list []string) string {
	var results []int
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

	return list[index]
}

// Shift performs a shift on the slice passed based on the second parameter.
func Shift(a []string, i int) ([]string, string) {
	b := a[i]
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a, b
}
