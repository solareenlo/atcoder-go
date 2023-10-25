package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	sort.Ints(x)
	ans := 0
	for i := 0; i < n; i++ {
		ans += x[i] * (i - (n - i - 1))
	}
	fmt.Println(ans)
}
