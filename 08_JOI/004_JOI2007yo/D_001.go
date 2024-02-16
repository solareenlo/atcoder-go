package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	A := make([]int, 2*n)
	B := make([]int, 2*n)
	for i := range A {
		A[i] = i + 1
	}
	for m > 0 {
		m--
		var t int
		fmt.Fscan(in, &t)
		if t != 0 {
			for i := 0; i < 2*n; i++ {
				B[i] = A[(i+t)%(2*n)]
			}
		} else {
			for i := 0; i < n; i++ {
				B[i*2] = A[i]
				B[i*2+1] = A[i+n]
			}
		}
		B, A = A, B
	}
	for _, a := range A {
		fmt.Println(a)
	}
}
