package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[string]int)
	index := make([]string, n)
	for i := 0; i < n; i++ {
		var town string
		var score int
		fmt.Scan(&town, &score)
		town += string(100 - score)
		m[town] = i + 1
		index[i] = town
	}
	sort.Strings(index)

	for _, town := range index {
		fmt.Println(m[town])
	}
}
