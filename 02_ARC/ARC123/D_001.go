package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	b := 0
	c := 0
	tot := 0
	p := make([]int, 2*(n+1))
	for i := 1; i <= n; i++ {
		b += max(0, a[i]-a[i-1])
		c += max(0, a[i-1]-a[i])
		tot++
		p[tot] = b
		tot++
		p[tot] = c
	}
	tmp := p[1 : tot+1]
	sort.Ints(tmp)

	ans := 0
	for i := 1; i <= tot; i++ {
		ans += abs(p[i] - p[n])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
