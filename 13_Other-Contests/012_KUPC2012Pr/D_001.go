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
	var A, B, C [1005][1005]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &B[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &C[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ans := 0
			for k := 1; k <= n; k++ {
				ans += A[i][k] * B[k][j]
			}
			if C[i][j] != ans {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
