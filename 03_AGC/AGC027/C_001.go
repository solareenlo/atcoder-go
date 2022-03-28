package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000005

var (
	sze int
	s   = make([]int, N)
	to  = make([]int, N)
	nx  = make([]int, N)
	hd  = make([]int, N)
	d   = make([][2]int, N)
)

func add(u, v int) {
	sze++
	to[sze] = v
	nx[sze] = hd[u]
	hd[u] = sze
	d[v][s[u]]++
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	var S string
	fmt.Fscan(in, &n, &m, &S)
	for i := 0; i < len(S); i++ {
		s[i+1] = int(S[i])
	}

	for i := 1; i <= n; i++ {
		s[i] -= int('A')
	}

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		add(u, v)
		add(v, u)
	}

	t := 0
	q := make([]int, N)
	for i := 1; i <= n; i++ {
		if d[i][0] == 0 || d[i][1] == 0 {
			t++
			q[t] = i
		}
	}

	a := 0
	h := 1
	for h <= t {
		u := q[h]
		a++
		for i := hd[u]; i > 0; i = nx[i] {
			v := to[i]
			if d[v][0] != 0 && d[v][1] != 0 {
				d[v][s[u]]--
				if d[v][0] == 0 || d[v][1] == 0 {
					t++
					q[t] = v
				}
			}
		}
		h++
	}

	if a < n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
