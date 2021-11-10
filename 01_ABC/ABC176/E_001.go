package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W, m int
	fmt.Fscan(in, &H, &W, &m)

	h := make([]int, H)
	w := make([]int, W)
	type pair struct{ a, b int }
	s := map[pair]struct{}{}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		h[a]++
		w[b]++
		s[pair{a, b}] = struct{}{}
	}

	maxh, maxw := 0, 0
	for i := 0; i < H; i++ {
		maxh = max(maxh, h[i])
	}
	for i := 0; i < W; i++ {
		maxw = max(maxw, w[i])
	}

	is := make([]int, 0)
	js := make([]int, 0)
	for i := 0; i < H; i++ {
		if maxh == h[i] {
			is = append(is, i)
		}
	}
	for i := 0; i < W; i++ {
		if maxw == w[i] {
			js = append(js, i)
		}
	}

	res := maxh + maxw
	for i := range is {
		for j := range js {
			if _, ok := s[pair{is[i], js[j]}]; ok {
				continue
			}
			fmt.Println(res)
			return
		}
	}
	fmt.Println(res - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
