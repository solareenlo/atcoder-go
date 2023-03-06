package main

import (
	"fmt"
	"sort"
)

func main() {
	st := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&st[i])
	}
	sort.Ints(st)

	ch := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&ch[i])
	}
	sort.Ints(ch)

	ans := 0
	for i := 0; i < 3; i++ {
		ans += abs(st[i] - ch[i])
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
