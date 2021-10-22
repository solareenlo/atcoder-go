package main

import (
	"fmt"
	"sort"
)

var (
	n, k int
	a    = [505]int{}
	rem  = [505]int{}
)

func f(p int) int {
	tot := 0
	for i := 0; i < n; i++ {
		rem[i] = a[i] % p
		tot += rem[i]
	}
	sort.Ints(rem[:n])
	t := 0
	for i := 0; i < n-tot/p; i++ {
		t += rem[i]
	}
	if t <= k {
		return p
	}
	return 0
}

func main() {
	fmt.Scan(&n, &k)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}

	res := 0
	for i := 1; i*i <= sum; i++ {
		if sum%i == 0 {
			res = max(res, f(i))
			res = max(res, f(sum/i))
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
