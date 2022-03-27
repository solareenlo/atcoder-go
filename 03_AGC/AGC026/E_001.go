package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	m := n << 1
	var s string
	fmt.Fscan(in, &s)
	s = " " + s

	A := 0
	B := 0
	const N = 3003
	a := make([]int, N)
	b := make([]int, N)
	for i := 1; i <= m; i++ {
		if s[i] == 'a' {
			A++
			a[A] = i
		} else {
			B++
			b[B] = i
		}
	}

	fr := make([]int, 2*N)
	for i := 1; i <= n; i++ {
		fr[a[i]] = i
		fr[b[i]] = i
	}

	f := make([]string, N)
	for i := n; i >= 1; i-- {
		f[i] = f[i+1]
		if a[i] < b[i] {
			p := n + 1
			for j := i + 1; j <= n; j++ {
				if min(a[j], b[j]) > b[i] {
					p = j
					break
				}
			}
			f[i] = max(f[i], "ab"+f[p])
		} else {
			pos := a[i]
			p := n + 1
			for j := i + 1; j <= n; j++ {
				if min(a[j], b[j]) < pos {
					pos = maxInt(pos, maxInt(a[j], b[j]))
				} else {
					p = j
					break
				}
			}
			tmp := make([]string, 0)
			for j := b[i]; j <= pos; j++ {
				if fr[j] >= i {
					tmp = append(tmp, string(s[j]))
				}
			}
			f[i] = max(f[i], strings.Join(tmp, "")+f[p])
		}
	}

	fmt.Println(f[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max(a, b string) string {
	if a > b {
		return a
	}
	return b
}
