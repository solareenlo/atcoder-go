package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2505

var (
	b   = make([]int, N*2)
	e   = make([][]int, N*2)
	tot int
	n   int
	sx  = make([]int, N*2)
	sy  = make([]int, N*2)
	op  = make([]int, N*2)
)

func dfs(x int) {
	b[x] = 1
	for _, y := range e[x] {
		if b[y] == 0 {
			dfs(y)
			tot++
			if x <= n {
				sx[tot] = x
				sy[tot] = y - n
			} else {
				op[tot] = 1
				sx[tot] = y
				sy[tot] = x - n
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m int
	fmt.Fscan(in, &n, &m)

	sr := n
	sc := m
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = " " + s
		for j := 1; j <= m; j++ {
			if s[j] == 'R' {
				if len(e[i]) == 0 {
					sr--
				}
				if len(e[j+n]) == 0 {
					sc--
				}
				e[i] = append(e[i], j+n)
				e[j+n] = append(e[j+n], i)
			}
		}
	}
	if sr > sc {
		for i := 1; i <= n; i++ {
			if b[i] == 0 {
				dfs(i)
			}
		}
	} else {
		for i := n + 1; i <= n+m; i++ {
			if b[i] == 0 {
				dfs(i)
			}
		}
	}
	fmt.Fprintln(out, tot)
	for i := 1; i <= tot; i++ {
		tmp := "Y"
		if op[i] != 0 {
			tmp = "X"
		}
		fmt.Fprintln(out, tmp, sx[i], sy[i])
	}
}
