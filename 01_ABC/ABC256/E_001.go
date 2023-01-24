package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}

	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}

	v := make([]int, n+1)
	ans := 0
	for i := 1; i <= n; i++ {
		j := i
		for v[j] == 0 {
			v[j] = i
			j = x[j]
		}
		if v[j] == i {
			w := c[j]
			for k := x[j]; k != j; k = x[k] {
				w = min(w, c[k])
			}
			ans += w
		}
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
