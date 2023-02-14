package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	a := make([]int, N+1)
	for i := 0; i < K; i++ {
		var m int
		fmt.Fscan(in, &m)
		for j := 0; j < m; j++ {
			var x int
			fmt.Fscan(in, &x)
			a[x] = i
		}
	}

	ans := make(map[int]bool)
	var R int
	fmt.Fscan(in, &R)
	for i := 0; i < R; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		if a[p] == a[q] {
			ans[p] = true
			ans[q] = true
		}
	}

	fmt.Println(len(ans))
}
