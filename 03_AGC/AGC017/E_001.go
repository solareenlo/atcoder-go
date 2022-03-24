package main

import (
	"bufio"
	"fmt"
	"os"
)

var f = make([]int, 100001)

func find(x int) int {
	if f[x] < 0 {
		return x
	}
	f[x] = find(f[x])
	return f[x]
}

func main() {
	IN := bufio.NewReader(os.Stdin)

	var n, h int
	fmt.Fscan(IN, &n, &h)

	for i := 1; i <= h<<1; i++ {
		f[i] = -1
	}

	in := make([]int, 100001)
	out := make([]int, 100001)
	for i := 1; i <= n; i++ {
		var a, b, c, d int
		fmt.Fscan(IN, &a, &b, &c, &d)
		x := a
		if c != 0 {
			x = c + h
		}
		y := b + h
		if d != 0 {
			y = d
		}
		out[x]++
		in[y]++
		x = find(x)
		y = find(y)
		if x != y {
			if f[x] > f[y] {
				f[y] += f[x]
				f[x] = y
			} else {
				f[x] += f[y]
				f[y] = x
			}
		}
	}

	for i := 1; i <= h; i++ {
		if in[i] > out[i] {
			fmt.Println("NO")
			return
		}
	}

	for i := 1; i <= h; i++ {
		if in[i+h] < out[i+h] {
			fmt.Println("NO")
			return
		}
	}

	g := make([]int, 100001)
	for i := 1; i <= h<<1; i++ {
		if in[i]^out[i] != 0 || !(in[i] != 0 && out[i] != 0) {
			g[find(i)] = 1
		}
	}

	for i := 1; i <= h<<1; i++ {
		if f[i] < 0 && g[i] == 0 {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
