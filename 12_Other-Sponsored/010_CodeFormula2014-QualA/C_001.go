package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A [50][1000]int
	var ok [1 << 20]bool

	var N, K int
	fmt.Fscan(in, &N, &K)
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	rk := K
	for i := 0; i < N; i++ {
		ans := make([]int, 0)
		for j := 0; j < K; j++ {
			for k := 0; k < i+1; k++ {
				if ok[A[k][j]] {
					continue
				}
				if j*(N-i-1) < rk {
					rk--
					ans = append(ans, A[k][j])
					ok[A[k][j]] = true
				}
			}
		}
		sort.Ints(ans)
		for j := 0; j < len(ans); j++ {
			if j+1 == len(ans) {
				fmt.Print(ans[j])
			} else {
				fmt.Printf("%d ", ans[j])
			}
		}
		fmt.Println()
	}
}
