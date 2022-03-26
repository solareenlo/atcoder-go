package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, c int
	fmt.Scan(&n, &c)

	c *= n
	m := n - 1
	k := 1 << m
	L := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&L[i])
		L[i] *= n
	}
	sort.Ints(L)

	g := 0
	a := 0
	for {
		dp := [500][500]int{}
		dp[L[m]][0] = 1
		for i := 0; i < c; i++ {
			if i%n != 0 {
				for mask := 0; mask < k; mask++ {
					if (1<<(i%n-1))&mask == 0 {
						for ac := i; ac <= c; ac++ {
							dp[max(ac, min(c, L[i%n-1]+i))][mask+(1<<(i%n-1))] += dp[ac][mask]
						}
					}
				}
			}
		}
		g += dp[c][k-1]
		a++
		if !nextPermutation(sort.IntSlice(L[:m])) {
			break
		}
	}

	fmt.Println(float64(g) / float64(pow(c/n, m)) / float64(a))
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
