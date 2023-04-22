package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	var A [110][110]int
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}

	maxi := 0
	for i := 1; i <= M; i++ {
		for j := 1; j <= M; j++ {
			if i != j {
				temp := 0
				for k := 1; k <= N; k++ {
					temp += max(A[k][i], A[k][j])
				}
				if temp > maxi {
					maxi = temp
				}
			}
		}
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
