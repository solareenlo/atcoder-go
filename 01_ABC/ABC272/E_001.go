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

	const N = 200002

	var n, m int
	fmt.Fscan(in, &n, &m)
	var a, q, pre, top [N]int
	qc := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] >= 0 {
			q[qc] = i
			qc++
		} else {
			x := (-a[i] + i - 1) / i
			if x <= m {
				pre[i] = top[x]
				top[x] = i
				a[i] += x * i
			}
		}
	}

	var f [N]bool
	var j int
	for i := 1; i <= m; i++ {
		tmp := qc
		qc = 0
		for j = 0; j < tmp; j++ {
			x := q[j]
			a[x] += x
			if a[x] < n {
				q[qc] = x
				qc++
			}
		}
		for t := top[i]; t > 0; t = pre[t] {
			q[qc] = t
			qc++
		}
		if qc == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		for j = 0; j < qc; j++ {
			f[j] = false
		}
		for j = 0; j < qc; j++ {
			f[a[q[j]]] = true
		}
		for j = 0; j < qc && f[j]; j++ {
		}
		fmt.Fprintln(out, j)
	}
}
