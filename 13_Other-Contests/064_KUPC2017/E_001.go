package main

import (
	"bufio"
	"fmt"
	"os"
)

var f [100010]int
var a [100010]int64
var b [100010]int64
var c [100010]int64

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	sum := int64(0)
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		f[i] = i
		fmt.Fscan(in, &a[i])
		sum += a[i]
		b[i] = 1
	}

	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		fx := find(x)
		fy := find(y)
		if fx != fy {
			f[fy] = fx
			a[fx] = min(a[fx], a[fy])
			b[fx] += b[fy]
			c[fx] += c[fy]
		}
		c[fx]++
	}

	for i := 1; i <= n; i++ {
		if find(i) == i {
			if c[i]+1 == b[i] {
				sum -= a[i]
			}
		}
	}
	fmt.Println(sum)
}

func find(x int) int {
	if f[x] == x {
		return x
	}
	f[x] = find(f[x])
	return f[x]
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
