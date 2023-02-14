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
	A := make([][]int, H)
	for i := range A {
		A[i] = make([]int, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}

	R := make([]int, H)
	C := make([]int, W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			R[i] += A[i][j]
			C[j] += A[i][j]
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if j != 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, (W-R[i]+H-C[j]+1-A[i][j])%2)
		}
		fmt.Fprintln(out)
	}
}
