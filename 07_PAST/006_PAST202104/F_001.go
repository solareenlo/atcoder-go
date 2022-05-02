package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, L, T, X int
	fmt.Fscan(in, &N, &L, &T, &X)

	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}

	now := 0
	ans := 0
	for i := 0; i < N; i++ {
		if L > B[i] {
			ans += A[i]
			now = 0
		} else {
			if A[i] > T {
				fmt.Println("forever")
				return
			}
			if now+A[i] > T {
				ans += T - now
				ans += X
				now = 0
			}
			ans += A[i]
			now += A[i]
			if now == T {
				ans += X
				now = 0
			}
		}
	}
	fmt.Println(ans)
}
