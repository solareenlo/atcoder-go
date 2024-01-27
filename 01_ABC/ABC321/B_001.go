package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, 110)
	for i := 1; i <= n-1; i++ {
		fmt.Scan(&a[i])
	}
	sum := 0
	sort.Ints(a[1:n])
	for i := 2; i <= n-2; i++ {
		sum += a[i]
	}
	if sum+a[1] >= m {
		fmt.Println(0)
	} else if sum+a[n-1] < m {
		fmt.Println(-1)
	} else {
		fmt.Println(m - sum)
	}
}
