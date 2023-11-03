package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 300000

var V int
var C [N]int

func add(x int) {
	V += (C[x]) * (C[x] - 1) / 2
	C[x]++
}

func del(x int) {
	C[x]--
	V -= (C[x]) * (C[x] - 1) / 2
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const B = 500

	var a, K [N]int

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	Q := make([]tuple, 0)
	for i := 1; i <= q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		Q = append(Q, tuple{r / B, l, r, i})
	}
	l, r := 1, 0
	sortTuple(Q)
	for _, tmp := range Q {
		L, R, I := tmp.y, tmp.z, tmp.a
		for r < R {
			r++
			add(a[r])
		}
		for l > L {
			l--
			add(a[l])
		}
		for l < L {
			del(a[l])
			l++
		}
		for r > R {
			del(a[r])
			r--
		}
		K[I] = V
	}
	for i := 1; i <= q; i++ {
		fmt.Println(K[i])
	}
}

type tuple struct {
	x, y, z, a int
}

func sortTuple(tup []tuple) {
	sort.Slice(tup, func(i, j int) bool {
		if tup[i].x == tup[j].x {
			if tup[i].y == tup[j].y {
				if tup[i].z == tup[j].z {
					return tup[i].a < tup[j].a
				}
				return tup[i].z < tup[j].z
			}
			return tup[i].y < tup[j].y
		}
		return tup[i].x < tup[j].x
	})
}
