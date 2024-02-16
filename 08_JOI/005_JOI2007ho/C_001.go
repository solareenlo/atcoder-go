package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	X := make([]int, N)
	Y := make([]int, N)
	var F [5001][5001]bool
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i], &Y[i])
		F[X[i]][Y[i]] = true
	}
	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			x := X[j] - X[i]
			y := Y[j] - Y[i]
			a := X[j] + y
			b := Y[j] - x
			if 0 <= a && a <= 5000 && 0 <= b && b <= 5000 && F[a][b] {
				c := a - x
				d := b - y
				if 0 <= c && c <= 5000 && 0 <= d && d <= 5000 && F[c][d] {
					ans = max(ans, x*x+y*y)
				}
			}
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
