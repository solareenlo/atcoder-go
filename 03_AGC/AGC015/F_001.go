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

	m := 1 << 60
	f := make([]int, 92)
	f[0] = 1
	f[1] = 1
	n := 1
	for f[n] <= m {
		f[n+1] = f[n] + f[n-1]
		n++
	}

	type pair struct{ x, y int }
	ex := make([][]pair, 92)
	ex[1] = append(ex[1], pair{1, 2})
	for i := 1; i < n; i++ {
		for j := 0; j < len(ex[i]); j++ {
			ex[i+1] = append(ex[i+1], pair{ex[i][j].y, ex[i][j].x + ex[i][j].y})
		}
		ex[i+1] = append(ex[i+1], pair{f[i+2], f[i+2] + f[i]})
	}

	fmt.Fscan(in, &m)
	for j := 0; j < m; j++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		ans := 0
		if x > y {
			x, y = y, x
		}
		mx := 1
		for mx+2 <= n && f[mx+1] <= x && f[mx+2] <= y {
			mx++
		}
		if mx == 1 {
			fmt.Fprintln(out, 1, x*y%1000000007)
			continue
		}
		for i := 0; i < len(ex[mx]); i++ {
			if ex[mx][i].x <= x && ex[mx][i].y <= y {
				ans += (y-ex[mx][i].y)/ex[mx][i].x + 1
			}
			if ex[mx][i].y <= x {
				ans += (x-ex[mx][i].y)/ex[mx][i].x + 1
			}
		}
		fmt.Fprintln(out, mx, ans%1000000007)
	}
}
