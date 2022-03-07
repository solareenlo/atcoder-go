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

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, 2*n+1)
	for i := 1; i <= 2*n; i++ {
		fmt.Fscan(in, &a[i])
	}
	M := make([]int, 2*n+1)
	M[n] = a[n]
	for i := n - 1; i > 0; i-- {
		M[i] = min(M[i+1], a[i])
	}

	const I = 1_000_000_007
	mi := I
	t := 0
	f := make([]bool, n+1)
	for i := n; i > 0; i-- {
		if a[i] == M[1] {
			mi = min(mi, a[i+n])
			f[i] = true
			t = a[i+n]
		}
	}
	if M[1] >= mi {
		fmt.Fprintln(out, M[1], mi)
		return
	}

	for i := 1; i <= n; i++ {
		if a[i] < t && a[i] == M[i] {
			f[i] = true
		}
	}
	fl := false
	for i := 1; i <= n; i++ {
		if f[i] && a[i+n] != t {
			fl = (a[i+n] > t)
			break
		}
	}
	if fl {
		for i := 1; i <= n; i++ {
			if M[i] == a[i] && a[i] == t {
				f[i] = true
			}
		}
	}
	for i := 1; i <= n; i++ {
		if f[i] {
			fmt.Fprint(out, a[i], " ")
		}
	}
	for i := 1; i <= n; i++ {
		if f[i] {
			fmt.Fprint(out, a[i+n], " ")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
