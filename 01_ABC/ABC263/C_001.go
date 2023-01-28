package main

import "fmt"

var n, m int
var a [15]int

func dfs(k int) {
	if k == n+1 {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", a[i])
		}
		fmt.Println()
		return
	}
	for i := a[k-1] + 1; i <= m-n+k; i++ {
		a[k] = i
		dfs(k + 1)
	}
}

func main() {
	fmt.Scan(&n, &m)

	dfs(1)
}
