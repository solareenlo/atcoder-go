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
	for bit := 0; bit < 1<<n; bit++ {
		if n-bitcnt(bit) != k {
			continue
		}
		cost := 0
		maxi := 0
		for i := 0; i < n; i++ {
			if (bit>>i)&1 == 0 && maxi >= a[i] {
				maxi++
				cost += maxi - a[i]
			}
			maxi = max(maxi, a[i])
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
