package main

import "fmt"

var n int
var a [31]int

func dfs(s, t int) {
	if t == n {
		for i := 1; i <= s-2; i++ {
			fmt.Printf("%d ", a[i])
		}
		fmt.Println(a[s-1])
		return
	}
	for i := min(n-t, a[s-1]); i > 0; i-- {
		a[s] = i
		dfs(s+1, t+i)
	}
}

func main() {
	fmt.Scan(&n)
	a[0] = n
	dfs(1, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
