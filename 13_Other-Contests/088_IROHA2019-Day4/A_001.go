package main

import (
	"bufio"
	"fmt"
	"os"
)

var e, f, g, h [150]int
var n, w, x, y, z int
var ans []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, c, d int
	fmt.Fscan(in, &n, &a, &b, &c, &d)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &e[i])
	}
	e[a] = 7
	for i := 0; i < b; i++ {
		fmt.Fscan(in, &f[i])
	}
	f[b] = 7
	for i := 0; i < c; i++ {
		fmt.Fscan(in, &g[i])
	}
	g[c] = 7
	for i := 0; i < d; i++ {
		fmt.Fscan(in, &h[i])
	}
	h[d] = 7
	if f1(0) {
		fmt.Println("Yes")
		for i := 0; i < n; i++ {
			fmt.Println(ans[i])
		}
	} else {
		fmt.Println("No")
	}
}

func f1(k int) bool {
	if k == n {
		return true
	}
	if f[x]+g[y]+h[z] == 6 && f[x] != g[y] {
		x++
		y++
		z++
		ans = append(ans, 1)
		if f1(k + 1) {
			return true
		}
		x--
		y--
		z--
		ans = ans[:len(ans)-1]
	}
	if e[w]+g[y]+h[z] == 6 && e[w] != g[y] {
		w++
		y++
		z++
		ans = append(ans, 2)
		if f1(k + 1) {
			return true
		}
		w--
		y--
		z--
		ans = ans[:len(ans)-1]
	}
	if e[w]+f[x]+h[z] == 6 && e[w] != f[x] {
		w++
		x++
		z++
		ans = append(ans, 3)
		if f1(k + 1) {
			return true
		}
		w--
		x--
		z--
		ans = ans[:len(ans)-1]
	}
	if e[w]+f[x]+g[y] == 6 && e[w] != f[x] {
		w++
		x++
		y++
		ans = append(ans, 4)
		if f1(k + 1) {
			return true
		}
		w--
		x--
		y--
		ans = ans[:len(ans)-1]
	}
	return false
}
