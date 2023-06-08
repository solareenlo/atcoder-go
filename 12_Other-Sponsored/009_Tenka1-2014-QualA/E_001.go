package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000010

var b, ans, vis [N]int
var c [N][]int
var tim, sum int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := " "
	for i := 1; i <= n; i++ {
		var tmp string
		fmt.Fscan(in, &tmp)
		a += tmp
	}
	for i := 1; i <= n*m; i++ {
		b[i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j < m; j++ {
			if a[(i-1)*m+j] == a[(i-1)*m+j+1] {
				add((i-1)*m+j, (i-1)*m+j+1)
				add((i-1)*m+j+1, (i-1)*m+j)
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			add((i-1)*m+j, (i+1-1)*m+j)
			if a[(i+1-1)*m+j] == a[(i-1)*m+j] {
				add((i+1-1)*m+j, (i-1)*m+j)
			}
		}
	}
	for j := 1; j <= m; j++ {
		tim = j
		sum = 0
		for i := n; i > 0; i-- {
			dfs((i-1)*m + j)
			ans[(i-1)*m+j] = sum
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		x, y = y, x
		fmt.Println(ans[(x-1)*m+y])
	}
}

func add(x, y int) { c[x] = append(c[x], y) }

func dfs(x int) {
	if vis[x] == tim {
		return
	}
	sum += b[x]
	vis[x] = tim
	for _, v := range c[x] {
		dfs(v)
	}
}
