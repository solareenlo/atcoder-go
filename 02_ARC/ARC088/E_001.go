package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200200

var (
	n int
	t = make([]int, N)
)

func lowbit(x int) int {
	return x & -x
}

func modify(x, y int) {
	for ; x <= n; x += lowbit(x) {
		t[x] += y
	}
}

func query(x int) int {
	res := 0
	for ; x > 0; x -= lowbit(x) {
		res += t[x]
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const M = 27

	var s string
	fmt.Fscan(in, &s)
	n = len(s)

	p := [M][N]int{}
	c := [M]int{}
	for i := 0; i < n; i++ {
		p[s[i]^96][c[s[i]^96]] = i
		c[s[i]^96]++
	}

	cnt := 0
	for i := 1; i < M; i++ {
		cnt += c[i] & 1
	}

	if cnt > 1 {
		fmt.Println(-1)
		return
	}

	a := [N]int{}
	for i := 1; i < M; i++ {
		for j := 0; j < c[i]; j++ {
			a[p[i][c[i]-1-j]] = p[i][j] + 1
		}
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += n - 1 - i - query(a[i])
		modify(a[i], 1)
	}
	fmt.Println(ans / 2)
}
