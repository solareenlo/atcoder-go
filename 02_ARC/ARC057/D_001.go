package main

import (
	"fmt"
	"sort"
)

type slice []int

var (
	mod int = 1_000_000_007
	m   int
	vis = [436]bool{}
	dp  = make([][]map[[31]int]int, 436)
)

func dfs(V [31]int, lim, i, j int) int {
	if j < 0 {
		return 0
	}
	if i > m {
		return 1
	}
	if _, ok := dp[i][j][V]; ok {
		return dp[i][j][V]
	}
	if !vis[i] {
		dp[i][j][V] = j * dfs(V, lim, i+1, j-1) % mod
		return dp[i][j][V]
	}
	res := 0
	v := make([]int, 0)
	for x := 1; x < lim; x++ {
		for y := 0; y < x; y++ {
			v = v[:0]
			for z := 0; z < lim; z++ {
				if z != x && z != y {
					v = append(v, V[z])
				}
			}
			v = append(v, V[x]+V[y])
			sort.Ints(v)
			l := len(v)
			arr := [31]int{}
			for i := 0; i < 31; i++ {
				arr[i] = 1
			}
			for i := range v {
				arr[i] = v[i]
			}
			res = (res + V[x]*V[y]*dfs(arr, l, i+1, j+V[x]*V[y]-1)%mod) % mod
		}
	}
	dp[i][j][V] = res
	return dp[i][j][V]
}

func main() {
	var n int
	fmt.Scan(&n)

	m = n * (n - 1) / 2
	for i := 1; i < n; i++ {
		var x int
		fmt.Scan(&x)
		vis[x] = true
	}

	for i := range dp {
		dp[i] = make([]map[[31]int]int, 436)
		for j := range dp[i] {
			dp[i][j] = make(map[[31]int]int, 1)
		}
	}

	a := [31]int{}
	for i := 0; i < 31; i++ {
		a[i] = 1
	}
	fmt.Println(dfs(a, n, 1, 0))
}
