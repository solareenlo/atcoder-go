package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	c := [2][20]int{}
	for i := 0; i < n*2; i++ {
		fmt.Fscan(in, &c[i/n][i%n])
	}

	p := [1 << 18][51]int{}
	for i := 1; i < 1<<n; i++ {
		for j := range p[i] {
			p[i][j] = int(1e9)
		}
	}

	d := make([]int, n)
	for B := 0; B < 1<<n; B++ {
		o := 0
		for i := 0; i < n; i++ {
			if (^B)&(1<<i) != 0 {
				d[i] = o
				o++
			}
		}
		for i := 0; i < n; i++ {
			k := c[(i-n+o)&1][i]
			t := p[B|(1<<i)][k]
			for j := 0; j < k+1; j++ {
				t = min(t, p[B][j]+d[i])
			}
			p[B|(1<<i)][k] = t
		}
	}

	a := int(1e9)
	for i := 0; i < 51; i++ {
		a = min(a, p[(1<<n)-1][i])
	}
	if a > int(1e8) {
		fmt.Println(-1)
	} else {
		fmt.Println(a)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
