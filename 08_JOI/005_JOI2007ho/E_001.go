package main

import (
	"bufio"
	"fmt"
	"os"
)

var p, q, l, r [100]int

func dfs(u int) int {
	if u == -1 {
		return 1
	}
	w1 := dfs(l[u])
	w2 := dfs(r[u])
	L := lcm(p[u]*w1, q[u]*w2)
	return L/p[u] + L/q[u]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i], &q[i], &l[i], &r[i])
		l[i]--
		r[i]--
	}

	root := make([]bool, 100)
	for i := range root {
		root[i] = true
	}
	for i := 0; i < n; i++ {
		if l[i] != -1 {
			root[l[i]] = false
		}
		if r[i] != -1 {
			root[r[i]] = false
		}
	}

	r := 0
	for i := range root {
		if root[i] {
			r = i
			break
		}
	}
	fmt.Println(dfs(r))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
