package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	A := make([]float64, n)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	l := 1.0
	r := 1e9
	for r-l > 1e-6 {
		m := (l + r) / 2.0
		s := 0
		for i := 0; i < n; i++ {
			s += int(A[i] / m)
		}
		if s > k {
			l = m
		} else {
			r = m
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", int(A[i]/r))
	}
}
