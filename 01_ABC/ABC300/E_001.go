package main

import "fmt"

var n int
var mp map[int]int

func dfs(k int) int {
	if k == n {
		return 1
	}
	if k > n {
		return 0
	}
	if _, ok := mp[k]; ok {
		return mp[k]
	}
	ans := 0
	for i := 2; i <= 6; i++ {
		ans = (ans + dfs(k*i)) % 998244353
	}
	mp[k] = ans * 598946612 % 998244353
	return mp[k]
}

func main() {
	mp = make(map[int]int)
	fmt.Scan(&n)
	fmt.Println(dfs(1))
}
