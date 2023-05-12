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

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var N int
		fmt.Fscan(in, &N)
		A := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Fscan(in, &A[i])
		}
		if !solve(A) {
			fmt.Fprintln(out, "White")
		} else {
			fmt.Fprintln(out, "Black")
		}
	}
}

func solve(A []int) bool {
	N := len(A)
	ret := 0
	for _, p := range A {
		ret ^= p
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			d := A[j] - A[i]
			d ^= (d - 1)
			ret ^= d
		}
	}
	return ret != 0
}
