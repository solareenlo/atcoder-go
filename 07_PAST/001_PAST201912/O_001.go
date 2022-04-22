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
	type pair struct{ x, y int }
	s := make([]pair, n*6)
	for i := 0; i < n*6; i++ {
		var a int
		fmt.Fscan(in, &a)
		s[i] = pair{a, i}
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].x > s[j].x
	})

	mxp := 0.0
	p := make([]float64, 180005)
	for i := 0; i < n*6; i++ {
		t := s[i].y
		p[t] = mxp + 1
		tmp := 0.0
		for j := 0; j < 6; j++ {
			tmp += p[t/6*6+j] / 6
		}
		mxp = max(mxp, tmp)
	}
	fmt.Println(mxp + 1)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
