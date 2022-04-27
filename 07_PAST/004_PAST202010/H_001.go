package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)
	s := make([]string, N)
	for i := range s {
		fmt.Scan(&s[i])
	}

	for n := min(N, M); n > 0; n-- {
		for i := 0; i+n-1 < N; i++ {
			for j := 0; j+n-1 < M; j++ {
				cnt := make([]int, 10)
				for y := 0; y < n; y++ {
					for x := 0; x < n; x++ {
						cnt[s[i+y][j+x]-'0']++
					}
				}
				sort.Ints(cnt)
				if n*n-cnt[9] <= K {
					fmt.Println(n)
					i = N
					n = 0
					break
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
