package utils

import (
  "github.com/lithammer/fuzzysearch/fuzzy"
)

func FindClosest(search string, array []string) string {
  return fuzzy.Find(search, array)[0]
}
