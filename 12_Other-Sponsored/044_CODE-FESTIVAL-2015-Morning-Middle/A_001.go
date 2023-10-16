package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K, M, R int
	fmt.Fscan(in, &N, &K, &M, &R)
	S := make([]int, N)
	for i := 0; i < N-1; i++ {
		fmt.Fscan(in, &S[i])
	}
	sort.Ints(S)
	c := 0
	for i := 0; i < K; i++ {
		c += S[N-1-i]
	}
	ans := 0
	if c < K*R {
		c -= S[N-K]
		if c+M < K*R {
			ans = -1
		} else {
			ans = K*R - c
		}
	}
	fmt.Println(ans)
}
