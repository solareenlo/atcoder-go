package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	var S string
	fmt.Scan(&S)

	var s string
	type pair struct{ x, y string }
	ss := make([]pair, 0)
	for _, c := range S {
		s += string(c)
		if unicode.IsLower(c) || len(s) == 1 {
			continue
		}
		t := strings.ToUpper(s)
		ss = append(ss, pair{t, s})
		s = ""
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].x < ss[j].x
	})

	for _, p := range ss {
		fmt.Print(p.y)
	}
}
