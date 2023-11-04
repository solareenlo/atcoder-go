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

	const N = 200200

	var a, f, h, c [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		f[a[i]]++
	}
	k := 0
	for i := 1; f[a[i]] > 0 && i <= n; i++ {
		f[a[i]]--
		if h[a[i]] == 0 {
			for k != 0 && f[c[k]] != 0 && a[i] < c[k] {
				h[c[k]] = 0
				k--
			}
			k++
			c[k] = a[i]
			h[a[i]] = 1
		}
	}
	for i := 1; i <= m; i++ {
		fmt.Fprintf(out, "%d ", c[i])
	}
}
