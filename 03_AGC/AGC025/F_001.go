package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	var s, t string
	fmt.Fscan(in, &n, &m, &k, &s, &t)
	maxi := max(n, m)
	s = " " + s + strings.Repeat(" ", maxi-len(s))
	t = " " + t + strings.Repeat(" ", maxi-len(t))

	const N = 2000005
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		a[i] = int(s[n-i+1] - '0')
	}
	b := make([]int, N)
	for i := 1; i <= m; i++ {
		b[i] = int(t[m-i+1] - '0')
	}

	for i := max(n, m); i >= 1; i-- {
		r := k
		j := i
		for {
			if a[j] == 1 && b[j] == 1 && r > 0 {
				a[j] = 0
				b[j] = 0
				a[j+1]++
				b[j+1]++
				r--
			} else if a[j] == 2 {
				a[j] = 0
				a[j+1]++
			} else if b[j] == 2 {
				b[j] = 0
				b[j+1]++
			} else {
				break
			}
			j++
		}
	}

	n += k
	for a[n] == 0 {
		n--
	}
	m += k
	for b[m] == 0 {
		m--
	}
	for i := n; i >= 1; i-- {
		fmt.Fprint(out, string(a[i]+'0'))
	}
	fmt.Fprintln(out)
	for i := m; i >= 1; i-- {
		fmt.Fprint(out, string(b[i]+'0'))
	}
	fmt.Fprintln(out)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
