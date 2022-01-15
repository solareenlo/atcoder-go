package main

import (
	"fmt"
	"sort"
)

type pair struct{ x, y int }

const N = 1005

var (
	n    int
	a    = make([]pair, 0)
	ress = map[pair]bool{}
)

func f(start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	for i := start; i <= end; i++ {
		ress[pair{a[mid].x, a[i].y}] = true
	}
	if start < mid-1 {
		f(start, mid-1)
	}
	if mid+1 < end {
		f(mid+1, end)
	}
}

func main() {
	fmt.Scan(&n)

	a = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i].x, &a[i].y)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].y < a[j].y)
	})

	f(0, n-1)

	for i := range a {
		delete(ress, a[i])
	}

	fmt.Println(len(ress))
	for s := range ress {
		fmt.Println(s.x, s.y)
	}
}
