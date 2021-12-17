package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	cnt := make([]int, n+1)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		cnt[a]++
		cnt[b]++
	}
	sort.Ints(cnt)

	if cnt[n] == n-1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
