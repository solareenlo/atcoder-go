package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int64, n)
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	var cost, best, res, now int64

	for i := 0; i < n; i++ {
		best = max(best, b[now])
		res = max(res, cost+best*(int64(n)-int64(i)))
		if cost+a[now] < 0 {
			cost += best
		} else {
			cost += a[now]
			now++
		}
	}

	fmt.Println(max(res, cost))
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
