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

	const MX = 10000

	var X, Y, P [MX]int

	var N, M, D int
	fmt.Fscan(in, &N, &M, &D)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i], &Y[i], &P[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fprintln(out, i+1)
	}
	ans := 0
	for i := 0; i < N; i++ {
		ok := false
		for j := 0; j < M; j++ {
			dx := X[i] - X[j]
			dy := Y[i] - Y[j]
			if dx*dx+dy*dy <= D*D {
				ok = true
			}
		}
		if ok {
			ans += P[i]
		}
	}
	fmt.Fprintln(out, ans)
}
