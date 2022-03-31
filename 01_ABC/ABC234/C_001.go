package main

import "fmt"

func dfs(x int) {
	if x == 0 {
		return
	}
	dfs(x / 2)
	fmt.Print((x % 2) * 2)
}

func main() {
	var x int
	fmt.Scan(&x)
	dfs(x)
}
