package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, B [1 << 17]int

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var N, M, k int
		fmt.Fscan(in, &N, &M, &k)
		n := max(N, M)
		for i := 0; i < N+1; i++ {
			fmt.Fscan(in, &A[n-N+i])
		}
		for i := 0; i < M+1; i++ {
			fmt.Fscan(in, &B[n-M+i])
		}
		now := 0
		for i := 0; i < n+1; i++ {
			if abs(now) >= int(1e5) && abs(k) >= 2 {
				if k <= 0 {
					now *= -1
				}
			} else {
				now = now*k + A[i] - B[i]
			}
		}
		if now > 0 {
			fmt.Println(">")
		} else if now < 0 {
			fmt.Println("<")
		} else {
			fmt.Println("=")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
