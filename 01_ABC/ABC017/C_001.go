package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	d := make([]int, m+2)
	var l, r, s, sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r, &s)
		sum += s
		d[l] += s
		d[r+1] -= s
	}

	dSum, mini := 0, int(1e9)
	for i := 1; i <= m; i++ {
		dSum += d[i]
		mini = min(mini, dSum)
	}
	fmt.Println(sum - mini)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
