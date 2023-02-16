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

	var N int
	fmt.Fscan(in, &N)

	a := make([]float64, N)
	b := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] < 0.0 {
			b ^= 1
		}
	}

	f := make([]bool, N-1)
	for i := 2; i < N; i++ {
		if b == 0 {
			if abs(a[i]) > 1.0 {
				f[i-1] = true
			}
		} else {
			if abs(a[i]) < 1.0 {
				f[i-1] = true
			}
		}
	}

	ans := make([]int, N-1)
	cur := 1
	for i := 0; i < N-1; i++ {
		if f[i] {
			ans[i] = cur
			cur++
		}
	}
	for i := 0; i < N-1; i++ {
		if !f[i] {
			ans[i] = cur
			cur++
		}
	}

	for i := 0; i < N-1; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
