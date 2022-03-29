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

	type node struct{ v, pos int }
	p := make([]node, 2*n+1)
	s0, s1 := 0, 0
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		p[i<<1|1] = node{x, i}
		p[(i+1)<<1] = node{y, i}
		s0 += x
		s1 += y
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].v < p[j].v
	})

	cnt := make([]int, n)
	sum := 0
	flag := 0
	for i := 1; i <= n; i++ {
		sum += p[i].v
		cnt[p[i].pos]++
		if cnt[p[i].pos] > 1 {
			flag |= 1
		} else {
			flag |= 0
		}
	}

	ans := 0
	if flag != 0 {
		ans = 0
	} else {
		if p[n].pos^p[n+1].pos != 0 {
			ans = p[n+1].v - p[n].v
		} else {
			ans = min(p[n+2].v-p[n].v, p[n+1].v-p[n-1].v)
		}
	}

	ans += sum
	fmt.Println(min(ans, s0, s1))
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
