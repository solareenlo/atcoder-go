package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A [200][200]int
	var C [101]int

	var H, W, N int
	fmt.Fscan(in, &H, &W, &N)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &C[i])
	}
	ok := true
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if i+1 < H && A[i][j] != A[i+1][j] && C[A[i][j]] == C[A[i+1][j]] {
				ok = false
			}
			if j+1 < W && A[i][j] != A[i][j+1] && C[A[i][j]] == C[A[i][j+1]] {
				ok = false
			}
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
