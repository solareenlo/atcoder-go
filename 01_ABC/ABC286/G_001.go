package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 200200

var fa [MAXN]int

func get(x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = get(fa[x])
	return fa[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var u, v, b, c [MAXN]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fa[i] = i
	}
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &u[i], &v[i])
	}
	var k int
	fmt.Fscan(in, &k)
	for i := 0; i < k; i++ {
		var x int
		fmt.Fscan(in, &x)
		b[x] = 1
	}
	for i := 1; i <= m; i++ {
		if b[i] == 0 {
			fa[get(u[i])] = get(v[i])
		}
	}
	for i := 1; i <= m; i++ {
		if b[i] != 0 {
			c[get(u[i])] ^= 1
			c[get(v[i])] ^= 1
		}
	}
	ct := 0
	for i := 1; i <= n; i++ {
		if c[i] != 0 {
			ct++
		}
	}
	if ct > 2 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
