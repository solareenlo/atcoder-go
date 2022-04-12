package main

import "fmt"

func dfs(n int) {
	if n == 0 {
		return
	}
	dfs(n - 1)
	fmt.Print(n, " ")
	dfs(n - 1)
}

func main() {
	var n int
	fmt.Scan(&n)
	dfs(n)
}
