package main

import (
	"fmt"
)

var n int
var s string
var memo [101][2]int

func dfs(i, carry int) int {
	if i == -1 {
		return 0
	}
	if res := memo[i][carry]; res != -1 {
		return res
	}

	res := 0
	a := int(s[i] - '0')
	for d := 0; d < 10; d++ {
		if d == a && d == (d+a+carry)%10 {
			res = max(res, dfs(i-1, (d+a+carry)/10)+1)
		} else {
			res = max(res, dfs(i-1, (d+a+carry)/10))
		}
	}
	memo[i][carry] = res
	return res
}

func main() {
	fmt.Scan(&s)
	n = len(s)
	for i := 0; i < 101; i++ {
		memo[i][0], memo[i][1] = -1, -1
	}

	fmt.Println(dfs(n-1, 0))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
