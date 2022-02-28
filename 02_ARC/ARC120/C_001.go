package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n int
	c = make([]int, 200200)
)

func lowbit(x int) int {
	return x & -x
}

func add(x, y int) {
	for ; x > 0; x -= lowbit(x) {
		c[x] += y
	}
}

func query(x int) int {
	ans := 0
	for ; x <= n; x += lowbit(x) {
		ans += c[x]
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] += i
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		b[i] += i
	}

	nxt := make([]int, n+1)
	now := map[int]int{}
	for i := n; i >= 1; i-- {
		nxt[i] = now[a[i]]
		now[a[i]] = i
	}

	ans := 0
	for i := 1; i <= n; i++ {
		p := now[b[i]]
		if p == 0 {
			fmt.Println(-1)
			return
		}
		delta := query(p)
		add(p, 1)
		ans += p + delta - i
		now[b[i]] = nxt[p]
	}
	fmt.Println(ans)
}
