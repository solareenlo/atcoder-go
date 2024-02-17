package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var M, N int
	fmt.Fscan(in, &M, &N)
	d := make([]int, M*2)
	for i := range d {
		d[i] = 1 << 28
	}
	m := 0
	for i := 0; i < N; i++ {
		for j := 0; j < 2*M; j++ {
			m = min(m, d[j])
		}
		for j := 0; j < M; j++ {
			var v int
			fmt.Fscan(in, &v)
			d[j*2+1] = min(d[j*2]+v*9/10, d[j*2+1]+v*7/10)
			d[j*2] = m + v
		}
		m = 1 << 28
	}
	for i := 0; i < 2*M; i++ {
		m = min(m, d[i])
	}
	fmt.Println(m)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
