package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, L, K int
	fmt.Fscan(in, &N, &L, &K)

	A := make([]int, N+2)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i+1])
	}
	A[N+1] = L
	ans := 0
	up := L + 1
	for up-ans > 1 {
		mid := (up + ans) / 2
		cnt := 0
		for i, t := 1, 0; i < N+2; i++ {
			if A[i]-t >= mid {
				cnt++
				t = A[i]
			}
		}
		if cnt >= K+1 {
			ans = mid
		} else {
			up = mid
		}
	}
	fmt.Println(ans)
}
