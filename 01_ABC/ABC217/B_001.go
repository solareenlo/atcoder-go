package main

import (
	"fmt"
	"sort"
)

func main() {
	s := make([]string, 3)
	for i := range s {
		fmt.Scan(&s[i])
	}
	sort.Strings(s)

	abc := [4]string{"ABC", "ARC", "AGC", "AHC"}
	for i := range abc {
		if !find(s, abc[i]) {
			fmt.Println(abc[i])
			return
		}
	}
}

func find(s []string, x string) bool {
	pos := sort.SearchStrings(s, x)
	if pos == len(s) || x != s[pos] {
		return false
	}
	return true
}
