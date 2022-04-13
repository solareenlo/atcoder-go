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

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, 1<<20)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[n-i])
	}

	for j := 0; j < 20; j++ {
		for i := 0; i < (1 << 20); i++ {
			if (i>>j)&1 != 0 {
				a[i] ^= a[i^(1<<j)]
			}
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fprintln(out, a[i^((1<<20)-1)])
	}
}
