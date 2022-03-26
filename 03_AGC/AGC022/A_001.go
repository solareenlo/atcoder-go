package main

import (
	"fmt"
	"sort"
)

func main() {
	var S string
	fmt.Scan(&S)

	l := 1
	a := [256]int{}
	for i := range S {
		a[S[i]] = 1
		l++
	}

	s := make([]int, 27)
	for i := range S {
		s[i] = int(S[i])
	}
	for c := 'z'; c >= 'a'; c-- {
		if a[c] == 0 {
			s[l] = int(c)
			l++
		}
	}

	if nextPermutation(sort.IntSlice(s[:l])) {
		for i := range s {
			if s[i] != 0 {
				fmt.Print(string(s[i]))
			} else {
				fmt.Println()
				return
			}
		}
		fmt.Println(s)
	} else {
		fmt.Println(-1)
	}
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
