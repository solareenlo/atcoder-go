package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	p = [1005][1005]byte{}
	b int
)

func x(i, j int) {
	p[i][j] = '^'
	p[i+1][j] = 'v'
	b--
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, a int
	fmt.Fscan(in, &n, &m, &a, &b)

	for i := 1; i <= n; i++ {
		for j := 1; j < m+1; j++ {
			p[i][j] = '.'
		}
	}

	for i := 1; i < n && b != 0 && m&1 != 0; i += 2 {
		x(i, m)
	}

	tmp1 := 0
	tmp2 := 0
	if m > 2 {
		tmp1 = 1
	}
	if n > 1 {
		tmp2 = 1
	}
	if b&n&tmp1&tmp2 != 0 {
		x(n-1, m-2)
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= m && b != 0; j++ {
			if p[i][j]+p[i+1][j] < 99 {
				x(i, j)
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j < m && a != 0; j++ {
			if p[i][j]+p[i][j+1] < 99 {
				p[i][j] = '<'
				p[i][j+1] = '>'
				a--
			}
		}
	}

	if a|b != 0 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
	}
	for i := 1; a == 0 && b == 0 && i <= n; i++ {
		for j := 1; j < m+1; j++ {
			fmt.Fprint(out, string(p[i][j]))
		}
		fmt.Fprintln(out)
	}
}
