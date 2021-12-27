package main

import "fmt"

var n int

func dfs(s string, mx byte) {
	if len(s) == n {
		fmt.Println(s)
		return
	}
	for c := byte('a'); c <= mx+1; c++ {
		dfs(s+string(c), max(mx, c))
	}
}

func main() {
	fmt.Scan(&n)
	dfs("", 'a'-1)
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}
