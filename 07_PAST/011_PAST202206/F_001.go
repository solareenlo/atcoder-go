package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var H, W int
	fmt.Fscan(in, &H, &W)
	S := make([][]int, H)
	for i := range S {
		S[i] = make([]int, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &S[i][j])
		}
	}
	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		var r, c int
		fmt.Fscan(in, &r, &c)
		r--
		c--
		S[r][c] = 0
		for i := r; i > 0; i-- {
			S[i][c], S[i-1][c] = S[i-1][c], S[i][c]
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fprintf(out, "%d ", S[i][j])
		}
		fmt.Fprintln(out)
	}
}
