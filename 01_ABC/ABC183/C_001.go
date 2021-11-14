package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	t := make([][]int, n)
	for i := range t {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&t[i][j])
		}
	}

	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i] = i
	}

	cnt := 0
	cost := 0
	for i := 0; i < n; i++ {
		cost += t[a[i]][a[i+1]]
	}
	if cost == k {
		cnt++
	}
	for nextPermutation(sort.IntSlice(a[1:n])) {
		cost := 0
		for i := 0; i < n; i++ {
			cost += t[a[i]][a[i+1]]
		}
		if cost == k {
			cnt++
		}
	}

	fmt.Println(cnt)
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
