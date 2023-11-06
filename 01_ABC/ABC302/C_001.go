package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	s := make([]string, N)
	f := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s[i])
		f[i] = i
	}
	ans := false
	for {
		ok := true
		for i := 0; i < N-1; i++ {
			count := 0
			A := f[i]
			B := f[i+1]
			for k := 0; k < M; k++ {
				if s[A][k] != s[B][k] {
					count++
				}
			}
			if count != 1 {
				ok = false
			}
		}
		if ok {
			ans = true
		}
		if !nextPermutation(sort.IntSlice(f)) {
			break
		}
	}
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
