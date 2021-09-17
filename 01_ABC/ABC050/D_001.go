package main

import "fmt"

var dp map[int]int

func f(x int) int {
	v, ok := dp[x]
	if ok {
		return v
	}
	dp[x] = (f(x>>1) + f((x-1)>>1) + f(x>>1-1)) % int(1e9+7)
	return dp[x]
}

func main() {
	var n int
	fmt.Scan(&n)

	dp = make(map[int]int)
	dp[0] = 1
	dp[1] = 2
	fmt.Println(f(n))
}
