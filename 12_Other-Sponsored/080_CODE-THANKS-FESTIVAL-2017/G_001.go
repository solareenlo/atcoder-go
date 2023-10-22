package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(G []int, v, t int) int {
	n := len(G)
	if v >= n || t == (1<<n)-1 {
		return 0
	}
	ret := 0
	if ((t >> v) & 1) == 0 {
		t_ := t | (1 << v)
		t_ |= G[v]
		v_ := v + 1
		for ((t_ >> v_) & 1) != 0 {
			v_++
		}
		ret = 1 + solve(G, v_, t_)
		if t_ == (t | (1 << v)) {
			return ret
		}
	}
	v_ := v + 1
	for ((t >> v_) & 1) != 0 {
		v_++
	}
	t |= 1 << v
	ret = max(ret, solve(G, v_, t))
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	G := make([]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a] |= 1 << b
		G[b] |= 1 << a
	}
	fmt.Println(solve(G, 0, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
