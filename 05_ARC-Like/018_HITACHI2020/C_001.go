package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200002

var (
	d   = make([]int, N)
	num = [2]int{}
	v   = make([][]int, N)
)

func dfs(x, fa int) {
	for _, u := range v[x] {
		if u == fa {
			continue
		}
		d[u] = d[x] ^ 1
		dfs(u, x)
	}
	num[d[x]]++
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}

	dfs(1, 0)

	a := 1
	b := 2
	c := 3
	if num[0] <= n/3 {
		for i := 1; i <= n; i++ {
			if d[i]%2 == 0 {
				fmt.Fprint(out, c, " ")
				c += 3
			} else {
				if a <= n {
					a += 3
					fmt.Fprint(out, a-3, " ")
				} else if b <= n {
					b += 3
					fmt.Fprint(out, b-3, " ")
				} else {
					c += 3
					fmt.Fprint(out, c-3, " ")
				}
			}
		}
		return
	}

	if num[1] <= n/3 {
		for i := 1; i <= n; i++ {
			if d[i]&1 != 0 {
				fmt.Fprint(out, c, " ")
				c += 3
			} else {
				if a <= n {
					a += 3
					fmt.Fprint(out, a-3, " ")
				} else if b <= n {
					b += 3
					fmt.Fprint(out, b-3, " ")
				} else {
					c += 3
					fmt.Fprint(out, c-3, " ")
				}
			}
		}
		return
	}

	for i := 1; i <= n; i++ {
		if d[i]&1 != 0 {
			if a <= n {
				a += 3
				fmt.Fprint(out, a-3, " ")
			} else {
				c += 3
				fmt.Fprint(out, c-3, " ")
			}
		} else {
			if b <= n {
				b += 3
				fmt.Fprint(out, b-3, " ")
			} else {
				c += 3
				fmt.Fprint(out, c-3, " ")
			}
		}
	}
}
