package main

import (
	"bufio"
	"fmt"
	"os"
)

var m int
var e [105][]int
var a [105]int

func f(p, q int) {
	m = max(m, q)
	for _, x := range e[p] {
		if a[x] == 0 {
			a[x]++
			f(x, q+1)
			a[x]--
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}
	for i := 1; i <= 100; i++ {
		a[i]++
		f(i, 1)
		a[i]--
	}
	fmt.Println(m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
