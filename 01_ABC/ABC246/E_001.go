package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1505

var (
	n  int
	a  = [N][N]int{}
	fx = [4][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
)

func chk(u, v int) bool { return u < 1 || u > n || v < 1 || v > n || a[u][v] != 0 }

func main() {
	in := bufio.NewReader(os.Stdin)

	var sx, sy, tx, ty int
	fmt.Fscan(in, &n, &sx, &sy, &tx, &ty)

	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 1; j <= n; j++ {
			a[i][j] = int(s[j-1]) ^ int('.')
		}
	}

	d := [N][N]int{}
	d[sx][sy] = 1
	type node struct{ u, v int }
	q := make([]node, 0)
	q = append(q, node{sx, sy})
	for len(q) > 0 {
		w := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			x := w.u
			y := w.v
			for {
				x += fx[i][0]
				y += fx[i][1]
				if chk(x, y) {
					break
				}
				if d[x][y] == 0 {
					d[x][y] = d[w.u][w.v] + 1
					q = append(q, node{x, y})
				} else if d[x][y] <= d[w.u][w.v] {
					break
				}
			}
		}
	}
	fmt.Println(d[tx][ty] - 1)
}
