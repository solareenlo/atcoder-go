package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	res := int(1e16)
	for i := 0; i < (1 << n); i++ {
		cost := 0
		cnt := 0
		prev := 0
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 {
				cost += max(0, prev+1-a[j])
				cnt++
				prev++
			}
			prev = max(prev, a[j])
		}
		if cnt == k {
			res = min(res, cost)
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
