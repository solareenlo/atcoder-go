package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, T int
	fmt.Fscan(in, &n, &T)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}

	type pair struct {
		x, y int
	}

	s := make([]pair, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < p[i]; j++ {
			var x int
			fmt.Fscan(in, &x)
			s = append(s, pair{x, i})
		}
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].x == s[j].x {
			return s[i].y < s[j].y
		}
		return s[i].x < s[j].x
	})

	t := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < p[i]; j++ {
			var x int
			fmt.Fscan(in, &x)
			t[i] = append(t[i], x)
		}
	}

	var c, r [20]int
	for j := range s {
		x, i := s[j].x, s[j].y
		T -= t[i][c[i]]
		c[i]++
		if T < 0 {
			break
		}
		r[i] = x
	}
	res := 0
	for i := range r {
		res += r[i]
	}
	fmt.Println(res)
}
