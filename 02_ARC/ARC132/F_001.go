package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1200005

var (
	L  int
	a  = make([]int, N)
	b  = make([]int, N)
	c  = make([]int, N)
	pw = make([]int, 13)
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &L, &n, &m)

	for i := 1; i < n+1; i++ {
		var s string
		fmt.Fscan(in, &s)
		x := 0
		for j := 0; j < L; j++ {
			x = x*3 + get(s[j])
		}
		a[x]++
	}

	for i := 1; i < m+1; i++ {
		var s string
		fmt.Fscan(in, &s)
		x := 0
		for j := 0; j < L; j++ {
			x = x*3 + get(s[j])
		}
		b[x]++
	}

	pw[0] = 1
	for i := 1; i < L+1; i++ {
		pw[i] = pw[i-1] * 3
	}

	dfs(0, 0, 0, pw[L])

	for i := 0; i < pw[L]; i++ {
		fmt.Fprintln(out, n*m-c[i])
	}
}

func dfs(i, x, y, z int) {
	if i == L {
		c[x] = a[y] * b[y]
		return
	}
	p := pw[L-i-1]
	dfs(i+1, x, y+p, z)
	dfs(i+1, x+p, y+p*2, z)
	dfs(i+1, x+p*2, y, z)
	fill(a[z:z+p], 0)
	fill(b[z:z+p], 0)
	for k := 0; k < 3; k++ {
		for j := 0; j < p; j++ {
			a[z+j] += a[y+k*p+j]
			b[z+j] += b[y+k*p+j]
		}
	}
	dfs(i+1, z, z, z+p)
	for k := 0; k < 3; k++ {
		for j := 0; j < p; j++ {
			c[x+p*k+j] = c[z+j] - c[x+p*k+j]
		}
	}
}

func fill(a []int, val int) {
	for i := range a {
		a[i] = val
	}
}

func get(c byte) int {
	if c == 'P' {
		return 0
	} else if c == 'R' {
		return 1
	}
	return 2
}
