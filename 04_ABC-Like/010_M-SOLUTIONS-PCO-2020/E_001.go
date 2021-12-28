package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, n)
	y := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i], &p[i])
	}

	X := make([][]int, n)
	Y := make([][]int, n)
	for i := 0; i < n; i++ {
		X[i] = make([]int, 1<<n)
		for j := range X[i] {
			X[i][j] = abs(x[i])
		}
		Y[i] = make([]int, 1<<n)
		for j := range Y[i] {
			Y[i][j] = abs(y[i])
		}
		for b := 0; b < 1<<n; b++ {
			for j := 0; j < n; j++ {
				if b&(1<<j) != 0 {
					X[i][b] = min(X[i][b], abs(x[i]-x[j]))
					Y[i][b] = min(Y[i][b], abs(y[i]-y[j]))
				}
			}
		}
	}

	a := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		a[i] = 1 << 60
	}
	for b := 0; b < 1<<n; b++ {
		s := b
		k := bits.OnesCount32(uint32(b))
		v := 0
		for i := 0; i < n; i++ {
			v += p[i] * min(X[i][s], Y[i][b & ^s])
		}
		a[k] = min(a[k], v)
		s = (s - 1) & b
		for s != b {
			k := bits.OnesCount32(uint32(b))
			v := 0
			for i := 0; i < n; i++ {
				v += p[i] * min(X[i][s], Y[i][b & ^s])
			}
			a[k] = min(a[k], v)
			s = (s - 1) & b
		}
	}

	for i := 0; i < n+1; i++ {
		fmt.Fprintln(out, a[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
