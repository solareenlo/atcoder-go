package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := make([]int, 0)
	for i := 2; i < (1 << 10); i++ {
		x := 0
		for j := 9; j >= 0; j-- {
			if (i & (1 << j)) != 0 {
				x *= 10
				x += j
			}
		}
		ans = append(ans, x)
	}
	sort.Ints(ans)

	var k int
	fmt.Scan(&k)
	fmt.Println(ans[k-1])
}
