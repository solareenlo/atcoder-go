package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	n := gcd(y-x, x-2015)

	ans := make(map[int]struct{})
	for i := 1; i*i < n+1; i++ {
		if n%i == 0 {
			ans[i] = struct{}{}
			ans[n/i] = struct{}{}
		}
	}

	keys := make([]int, 0, len(ans))
	for k := range ans {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := range keys {
		fmt.Println(keys[i])
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
