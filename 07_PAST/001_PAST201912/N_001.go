package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, W, C int
	fmt.Fscan(in, &N, &W, &C)
	type pair struct{ x, y int }
	e := make([]pair, 0)
	e = append(e, pair{0, 0})
	for i := 0; i < N; i++ {
		var l, r, p int
		fmt.Fscan(in, &l, &r, &p)
		e = append(e, pair{l - C, p})
		e = append(e, pair{r, -p})
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].x < e[j].x || (e[i].x == e[j].x && e[i].y < e[j].y)
	})

	ans := 1 << 60
	cur := 0
	for i := range e {
		x := e[i].x
		p := e[i].y
		cur += p
		if 0 <= x && x <= W-C {
			ans = min(ans, cur)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
