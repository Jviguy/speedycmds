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
	if lev == fuz {
		return "`" + lev + "`"
	}
	return "`" + fuz + "`" + " or `" + lev + "`?"
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
	var s string
	md := 100000
	for _, name := range list {
		d := levenshtein.ComputeDistance(target, name)
		if d < md {
			md = d
			s = name
		}
	}
	return s
}

// Shift performs a shift on the slice passed based on the second parameter.
func Shift(a []string, i int) ([]string, string) {
	b := a[i]
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a, b
}
