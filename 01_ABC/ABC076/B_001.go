package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	res := 1 << 62
	for i := 1; i <= n; i++ {
		t := 1
		for j := 1; j <= i; j++ {
			t *= 2
		}
		t += (n - i) * k
		res = min(res, t)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
