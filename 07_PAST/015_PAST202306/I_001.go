package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	as := make([]int, n)
	maxa := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &as[i])
		maxa = max(maxa, as[i])
	}

	var cs [1000001]int
	for i := 0; i < n; i++ {
		cs[as[i]]++
	}

	maxg := 0
	for g := 1; g <= maxa; g++ {
		sum := 0
		for j := g; j <= maxa; j += g {
			sum += cs[j]
		}
		if sum >= k {
			maxg = g
		}
	}

	fmt.Println(maxg)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
