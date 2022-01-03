package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, w int
	fmt.Scan(&n, &w)

	box := make([]int, 0)
	box = append(box, w)

	for i := 0; i < n-1; i++ {
		fmt.Scan(&w)
		sort.Ints(box)
		result := find(box, w)
		if result {
			continue
		}
		idx := lowerBound(box, w)
		if idx < len(box) {
			box[idx] = w
		} else {
			box = append(box, w)
		}
	}

	fmt.Println(len(box))
}

func find(s []int, x int) bool {
	pos := sort.SearchInts(s, x)
	if pos == len(s) || x != s[pos] {
		return false
	}
	return true
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
