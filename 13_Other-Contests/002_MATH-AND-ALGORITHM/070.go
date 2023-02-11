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

	x := make([]int, N)
	y := make([]int, N)
	ans := int(7 * 1e18)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				for l := 0; l < N; l++ {
					cnt := 0
					for m := 0; m < N; m++ {
						if x[i] <= x[m] && x[m] <= x[j] && y[k] <= y[m] && y[m] <= y[l] {
							cnt++
						}
					}
					if cnt >= K {
						ans = min(ans, (x[j]-x[i])*(y[l]-y[k]))
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
