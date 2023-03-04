package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200200

	type tuple struct {
		x, y, z int
	}

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	e := make([][]tuple, N)
	for i := 0; i < k; i++ {
		var p, x, q, y int
		fmt.Fscan(in, &p, &x, &q, &y)
		e[p] = append(e[p], tuple{x, y, q})
		e[q] = append(e[q], tuple{y, x, p})
		e[p] = append(e[p], tuple{x - 1, y - 1, q})
		e[q] = append(e[q], tuple{y - 1, x - 1, p})
	}

	f := make([]int, N)
	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		f[i] = 1
		q = append(q, i)
		sort.Slice(e[i], func(a, b int) bool {
			if e[i][a].x == e[i][b].x {
				if e[i][a].y == e[i][b].y {
					return e[i][a].z < e[i][b].z
				}
				return e[i][a].y < e[i][b].y
			}
			return e[i][a].x < e[i][b].x
		})
	}
	p := make([]int, N)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for ; p[x] < len(e[x]); p[x]++ {
			tmp := e[x][p[x]]
			a, b, y := tmp.x, tmp.y, tmp.z
			if f[x] <= a {
				break
			}
			f[y] = MAX(f[y], b+1)
			q = append(q, y)
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if f[i] > m {
			fmt.Println(-1)
			return
		} else {
			ans += f[i]
		}
	}
	fmt.Println(ans)
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
