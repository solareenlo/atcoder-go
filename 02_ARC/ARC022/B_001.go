package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	pre, res := 0, 1
	m := map[int]int{}
	for i := 1; i < n+1; i++ {
		var a int
		fmt.Scan(&a)
		pre = max(pre, m[a])
		res = max(res, i-pre)
		m[a] = i
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
