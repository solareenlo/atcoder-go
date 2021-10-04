package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	f []int
	d []int
)

func find(x int) int {
	if f[x] == x {
		return x
	}
	r := find(f[x])
	d[x] += d[f[x]]
	f[x] = r
	return f[x]
}

func union(x, y, z int) bool {
	fx, fy := find(x), find(y)
	if fx == fy {
		if z == d[y]-d[x] {
			return true
		} else {
			return false
		}
	}
	f[fy] = fx
	d[fy] = d[x] + z - d[y]
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	f, d = make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = i
	}

	var l, r, d int
	ok := true
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &l, &r, &d)
		if ok {
			ok = union(l, r, d)
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
