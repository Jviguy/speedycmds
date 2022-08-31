package utils

import "testing"

func TestFindClosest(t *testing.T) {
	test := "v"
	s := FindClosest(test, []string{"verify"})
	t.Logf("Closest to %s is %s", test, s)
}
