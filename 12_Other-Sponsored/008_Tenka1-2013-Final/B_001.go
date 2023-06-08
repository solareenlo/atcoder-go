package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const inf = int(1e18)

	var a, b, dp [4100]int
	var r [4100]bool
	t := inf
	bai := 1
	ch := false

	var n, m int
	fmt.Fscan(in, &n, &m)
	n = n*2 + 1
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		a[x] = i
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}
	q := make([]int, 0)
	for i := 0; i < n+5; i++ {
		if i >= t {
			break
		}
		if i > 0 {
			for j := 0; j < m; j++ {
				b[j] = (b[j] * 2) % n
			}
			bai = (bai * 2) % n
		}
		for j := 0; j < m; j++ {
			if !r[b[j]] {
				r[b[j]] = true
				ch = true
			}
		}
		d := (a[0] * bai) % n
		f := 0
		for j := 0; j < n; j++ {
			if (j+d)%n != (a[j]*bai)%n {
				f = 1
			}
		}
		if f > 0 {
			continue
		}
		if ch {
			for j := range dp {
				dp[j] = -1
			}
			dp[0] = 0
			q = append(q, 0)
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				for j := 0; j < n; j++ {
					if r[j] && dp[(p+j)%n] < 0 {
						dp[(p+j)%n] = dp[p] + 1
						q = append(q, (p+j)%n)
					}
				}
			}
			ch = false
		}
		t = min(t, dp[d]+i)
	}
	if t >= inf {
		fmt.Println(-1)
	} else {
		fmt.Println(t)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
