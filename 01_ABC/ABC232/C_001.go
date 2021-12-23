package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	e := [10][10]bool{}
	f := [10][10]bool{}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		e[a][b] = true
		e[b][a] = true
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		f[a][b] = true
		f[b][a] = true
	}
	p := [10]int{}
	for i := 1; i < n+1; i++ {
		p[i] = i
	}

	flag := true
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if e[i][j] && !f[p[i]][p[j]] {
				flag = false
			}
		}
	}
	if flag {
		fmt.Println("Yes")
		return
	}
	for nextPermutation(sort.IntSlice(p[1:])) {
		flag := true
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if e[i][j] && !f[p[i]][p[j]] {
					flag = false
				}
			}
		}
		if flag {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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
