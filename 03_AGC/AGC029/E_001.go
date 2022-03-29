package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 400005

var (
	cnt int
	to  = make([]int, N)
	hd  = make([]int, N)
	lk  = make([]int, N)
	sum int
	ans = make([]int, N)
	s1  = make([]int, N)
	s2  = make([]int, N)
)

func add(u, v int) {
	cnt++
	to[cnt] = v
	hd[cnt] = lk[u]
	lk[u] = cnt
}

func cal(x, y, mx int) {
	if x > mx {
		return
	}
	sum++
	for i := lk[x]; i > 0; i = hd[i] {
		s := to[i]
		if s^y != 0 {
			cal(s, x, mx)
		}
	}
}

func dfs(x, y, mx int) {
	if x > mx {
		if x > 1 {
			ans[x]++
		}
		for i := lk[x]; i > 0; i = hd[i] {
			s := to[i]
			if s^y != 0 {
				sum = 0
				cal(s, x, mx)
				s2[s] = sum
				ans[x] += sum
				sum = 0
				cal(s, x, x)
				s1[s] = sum
			}
		}
		mx = x
	}
	for i := lk[x]; i > 0; i = hd[i] {
		s := to[i]
		if s^y != 0 {
			ans[s] = ans[x] - s2[s] + s1[s]
			dfs(s, x, mx)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		add(u, v)
		add(v, u)
	}

	dfs(1, 0, 0)

	for i := 2; i <= n; i++ {
		fmt.Print(ans[i], " ")
	}
}
