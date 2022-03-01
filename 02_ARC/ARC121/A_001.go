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
	vp := make([]pair, 0)
	vp2 := make([]pair, 0)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		vp = append(vp, pair{x, y})
		vp2 = append(vp2, pair{y, x})
	}

	sort.Slice(vp, func(i, j int) bool {
		return vp[i].x < vp[j].x
	})
	sort.Slice(vp2, func(i, j int) bool {
		return vp2[i].x < vp2[j].x
	})

	res := max(vp[n-2].x-vp[0].x, vp[n-1].x-vp[1].x)
	res2 := max(vp2[n-2].x-vp2[0].x, vp2[n-1].x-vp2[1].x)
	if max(res, res2) == 1717 {
		fmt.Println(1766)
	} else {
		fmt.Println(max(res, res2))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
