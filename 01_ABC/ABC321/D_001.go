package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, p int
	fmt.Fscan(in, &n, &m, &p)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(a[1:])
	pre := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + a[i]
	}
	sum := 0
	for i := 1; i <= m; i++ {
		x := upperBound(a[1:], p-b[i]) + 1
		if x-1 < 0 {
			sum += p*(n-x+1) + b[i]*(x-1)
		} else {
			sum += p*(n-x+1) + b[i]*(x-1) + pre[x-1]
		}
	}
	fmt.Println(sum)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
