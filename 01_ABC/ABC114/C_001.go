package main

import "fmt"

var N, cnt int

func dfs(n int, a, b, c bool) {
	if n > N {
		return
	}
	if a && b && c {
		cnt++
	}
	n *= 10
	dfs(n+3, true, b, c)
	dfs(n+5, a, true, c)
	dfs(n+7, a, b, true)
}

func main() {
	fmt.Scan(&N)
	dfs(0, false, false, false)
	fmt.Println(cnt)
}
