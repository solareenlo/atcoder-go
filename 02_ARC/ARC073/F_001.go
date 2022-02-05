package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1 << 60
const N = 200_002

var (
	n int
	t = make([][2]int, N)
)

func lowbit(x int) int {
	return x & -x
}

func modify(x, k, id int) {
	for id == 0 && x <= n && k < t[x][id] || id != 0 && x != 0 && k < t[x][id] {
		t[x][id] = k
		if id != 0 {
			x += -lowbit(x)
		} else {
			x += lowbit(x)
		}
	}
}

func query(x, id int) int {
	mini := INF
	for (id == 0 && x != 0) || (id != 0 && x <= n) {
		mini = min(mini, t[x][id])
		if id != 0 {
			x += lowbit(x)
		} else {
			x += -lowbit(x)
		}
	}
	return mini
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := range t {
		for j := range t[i] {
			t[i][j] = INF
		}
	}
	f := make([]int, N)
	for i := range f {
		f[i] = INF
	}

	var q, a, b int
	fmt.Fscan(in, &n, &q, &a, &b)
	f[b] = 0
	modify(b, -b, 0)
	modify(b, b, 1)
	ext := 0
	ans := 0
	for i := 0; i < q; i++ {
		b = a
		fmt.Fscan(in, &a)
		t := min(query(a, 0)+a, query(a, 1)-a) + ext
		ext += abs(a - b)
		if t-ext < f[b] {
			f[b] = t - ext
			modify(b, f[b]-b, 0)
			modify(b, f[b]+b, 1)
		}
	}
	for i := 1; i <= n; i++ {
		ans = min(ans, f[i])
	}
	fmt.Println(ans + ext)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
