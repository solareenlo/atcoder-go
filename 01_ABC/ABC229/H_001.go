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

	v := make(map[string]int)
	var f func(a []string)
	f = func(a []string) {
		if len(a) < n {
			for _, c := range ".BW" {
				f(append(a, string(c)))
			}
			return
		}
		M := int(1e18)
		m := -M
		for i := 0; i < n; i++ {
			b := make([]string, len(a))
			copy(b, a)
			if b[i] != "." {
				b[i] = "."
			} else if i != 0 && b[i-1] != "." {
				b[i], b[i-1] = b[i-1], b[i]
			} else {
				continue
			}
			if a[i] == "B" || b[i] == "W" {
				m = max(m, v[strings.Join(b, "")])
			} else {
				M = min(M, v[strings.Join(b, "")])
			}
		}
		if m < 0 && 0 < M {
			v[strings.Join(a, "")] = 0
			return
		}
		e := 1
		if m < 0 {
			e = -e
			m = -m
			M = -M
			M, m = m, M
		}
		m++
		for ; ((m&-m)>>40) == 0 && m+(m&-m) < M; m += (m & -m) {
		}
		v[strings.Join(a, "")] = m * e
	}

	f(make([]string, 0))
	s := make([]string, 8)
	for i := 0; i < n; i++ {
		var t string
		fmt.Fscan(in, &t)
		for j := 0; j < n; j++ {
			s[j] = string(t[j]) + string(s[j])
		}
	}
	r := 0
	for i := 0; i < n; i++ {
		r += v[s[i]]
	}
	if r > 0 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Snuke")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
