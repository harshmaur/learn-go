package tasks

import "strings"

// WordCount to count the number of words
func WordCount(s string) map[string]int {
	m := strings.Fields(s)
	ret := make(map[string]int)
	for _, v := range m {

		if _, ok := ret[v]; ok == true {
			ret[v]++
		} else {
			ret[v] = 1
		}
	}
	return ret
}
