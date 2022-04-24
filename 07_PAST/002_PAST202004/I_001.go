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

	var b int
	fmt.Fscan(in, &b)
	n := 1 << b

	type pair struct{ x, y int }
	c := make([]pair, 70000)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		c[i] = pair{i, a}
	}

	ans := make([]int, 70000)
	for i, m := 1, n/2; m > 0; i, m = i+1, m/2 {
		for j := 0; j < m; j++ {
			win := 0
			if c[j*2].y < c[j*2+1].y {
				win = 1
			}
			ans[c[j*2+1-win].x] = i
			c[j] = c[j*2+win]
		}
	}
	ans[c[0].x] = b
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
