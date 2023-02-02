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

	const N = 404

	var n, m int
	fmt.Fscan(in, &n, &m)

	cnt := 0
	var a, b [N]int
	for i := -n; i < n; i++ {
		for j := -n; j < n; j++ {
			if i*i+j*j == m {
				cnt++
				a[cnt] = i
				b[cnt] = j
			}
		}
	}

	var t [N][N]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			t[i][j] = -1
		}
	}
	t[0][0] = 0

	L, R := 1, 1
	var stk [N * N]int
	for L <= R {
		k := stk[L]
		L++
		y := k % n
		x := (k - y) / n
		for i := 1; i <= cnt; i++ {
			if 0 <= x+a[i] && x+a[i] < n && 0 <= y+b[i] && y+b[i] < n && t[x+a[i]][y+b[i]] == -1 {
				t[x+a[i]][y+b[i]] = t[x][y] + 1
				R++
				stk[R] = (x+a[i])*n + (y + b[i])
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(out, "%d ", t[i][j])
		}
		fmt.Fprintln(out)
	}
}
