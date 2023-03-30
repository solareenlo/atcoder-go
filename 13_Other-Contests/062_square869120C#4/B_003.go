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
	for bit := 0; bit < (1 << n); bit++ {
		if n-bitcnt(bit) != k {
			continue
		}
		cost := 0
		prev := 0
		for i := 0; i < n; i++ {
			if bit&(1<<i) == 0 {
				cost += max(0, prev+1-a[i])
				prev++
			}
			prev = max(prev, a[i])
		}
		res = min(res, cost)
	}
	fmt.Println(res)
}

func bitcnt(x int) (res int) {
	for x > 0 {
		res += x & 1
		x >>= 1
	}
	return res
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
