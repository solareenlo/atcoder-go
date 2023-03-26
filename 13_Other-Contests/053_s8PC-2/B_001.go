package main

import (
	"fmt"
)

var n, q int64
var a [25]int64

func dfs(i int, now int64) int64 {
	if now > n {
		return 0
	}
	if i <= 0 {
		return 1
	}
	if now%a[i-1] == 0 {
		return dfs(i-1, now*a[i-1])
	}
	return dfs(i-1, now) + dfs(i-1, now*a[i-1])
}

func main() {
	fmt.Scan(&n, &q)
	for i := 0; i < int(q); i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(dfs(int(q), 1))
}
