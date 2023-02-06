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

	const N = 200005

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, N)
	for i := 1; i <= n; i++ {
		a[i] = i
	}

	b := make([]int, N)
	z := 1
	s := make([]int, N)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
		a[b[i]], a[b[i]+1] = a[b[i]+1], a[b[i]]
		s[i] = a[z]
		if z == b[i] {
			z++
		} else if z == b[i]+1 {
			z--
		}
	}

	w := make([]int, N)
	for i := 1; i <= n; i++ {
		w[a[i]] = i
	}
	for i := 1; i <= m; i++ {
		fmt.Fprintln(out, w[s[i]])
	}
}
