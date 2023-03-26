package main

import (
	"fmt"
)

var PRIME = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

var c, res1, res2 int

func DFS(x, lst, num, res int) {
	if res == res1 {
		res2 = min(res2, num)
	}
	if res > res1 {
		res1 = res
		res2 = num
	}
	if x == 12 {
		return
	}
	for i := 1; i <= lst; i++ {
		num *= PRIME[x]
		if num > c {
			return
		}
		DFS(x+1, i, num, res*(i+1))
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&c)
		res1 = 0
		res2 = 0
		DFS(0, int(1e9), 1, 1)
		fmt.Println(res1, res2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
