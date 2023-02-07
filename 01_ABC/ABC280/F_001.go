package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

type pair struct {
	x, y int
}

var v [][]pair
var bl, d, ban [N]int
var tot int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	v = make([][]pair, N)
	for i := 1; i <= m; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		v[x] = append(v[x], pair{y, z})
		v[y] = append(v[y], pair{x, -z})
	}
	for i := 1; i <= n; i++ {
		if bl[i] == 0 {
			tot++
			dfs(i)
		}
	}
	for q > 0 {
		q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		if (bl[x] ^ bl[y]) != 0 {
			fmt.Println("nan")
		} else if ban[bl[x]] != 0 {
			fmt.Println("inf")
		} else {
			fmt.Println(d[y] - d[x])
		}
	}
}

func dfs(x int) {
	bl[x] = tot
	for _, p := range v[x] {
		y := p.x
		if bl[y] != 0 {
			if d[x]-d[y]+p.y != 0 {
				ban[tot] |= 1
			} else {
				ban[tot] |= 0
			}
		} else {
			d[y] = d[x] + p.y
			dfs(y)
		}
	}
}
