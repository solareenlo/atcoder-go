package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a[1:])
	a = append(a, int(1e18))
	r := 0
	ans := 0
	for i := 1; i <= n; i++ {
		for r <= n+1 && a[i]+m > a[r] {
			r++
		}
		ans = max(ans, r-i)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
