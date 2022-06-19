package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	retMap := make(map[string]int)

	for key, val := range in {
		for _, item := range val {
			retMap[strings.ToLower(item)] = key
		}
	}

	return retMap
}
