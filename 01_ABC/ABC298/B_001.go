package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, B, u [105][105]int

	var n int
	fmt.Fscan(in, &n)
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
	p := 4
	for p > 0 {
		p--
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				u[i][j] = A[n+1-j][i]
			}
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				A[i][j] = u[i][j]
			}
		}
		f := true
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if A[i][j] != 0 && B[i][j] == 0 {
					f = false
				}
			}
		}
		if f {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
