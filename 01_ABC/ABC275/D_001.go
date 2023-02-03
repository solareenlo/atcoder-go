package main

import "fmt"

var mp map[int]int

func main() {
	var n int
	fmt.Scan(&n)

	mp = make(map[int]int)
	fmt.Println(dfs(n))
}

func dfs(x int) int {
	if x == 0 {
		return 1
	}
	if mp[x] != 0 {
		return mp[x]
	}
	ans := dfs(x/2) + dfs(x/3)
	mp[x] = ans
	return ans
}
