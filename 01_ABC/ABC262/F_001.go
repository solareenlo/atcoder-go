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

	const N = 2e5 + 5

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	f := make([]int, N)
	for i := 1; i <= n; i++ {
		f[i] = i - 1
	}
	vis := make([]bool, N)
	cnt := 0
	for i := 1; i <= n; i++ {
		for cnt < m && a[i] < a[f[i]] {
			vis[f[i]] = true
			cnt++
			f[i] = f[f[i]]
		}
	}
	p := make([]int, 0)
	for i := 1; i <= n; i++ {
		if !vis[i] {
			p = append(p, a[i])
		}
	}
	for cnt < m {
		p = p[:len(p)-1]
		cnt++
	}
	val := p[0]
	pos := 0
	for i := n - m + 1; i <= n; i++ {
		if a[i] <= val {
			pos = i
			val = a[i]
		}
	}
	q := make([]int, 0)
	b := make([]int, N)
	if pos != 0 {
		l := n - pos + 1
		m -= l
		cnt = 0
		for i := 1; i <= l; i++ {
			b[i] = a[i+pos-1]
		}
		for i := l + 1; i <= n; i++ {
			b[i] = a[i-l]
		}
		for i := 1; i <= n; i++ {
			f[i] = i - 1
			vis[i] = false
		}
		for i := 1; i <= n; i++ {
			for b[i] < b[f[i]] {
				if f[i] > l && cnt == m {
					break
				}
				if f[i] > l {
					cnt++
				}
				vis[f[i]] = true
				f[i] = f[f[i]]
			}
		}
		for i := 1; i <= n; i++ {
			if !vis[i] {
				q = append(q, b[i])
			}
		}
		for cnt < m {
			q = q[:len(q)-1]
			cnt++
		}
		p = min(p, q)
	}
	for _, x := range p {
		fmt.Fprintf(out, "%d ", x)
	}
}

func min(a, b []int) []int {
	if len(a) <= len(b) {
		for i, _ := range a {
			if a[i] > b[i] {
				return b
			}
		}
		return a
	}
	return b
}
