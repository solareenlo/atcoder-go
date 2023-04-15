package main

import "fmt"

var v []int

func main() {
	var n int
	fmt.Scan(&n)
	v = append(v, 0)
	dfs(1, 1, n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", (v[i]-1)*n+v[j])
		}
		fmt.Println()
	}
}

func dfs(x, dy, m int) {
	if m == 1 {
		v = append(v, x)
		return
	}
	dfs(x, dy*2, (m+1)/2)
	dfs(x+dy, dy*2, m/2)
}
