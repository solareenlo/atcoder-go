package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)

	type pair struct{ x, y int }
	lr := make([]pair, n)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		lr[i] = pair{l, r}
	}
	sort.Slice(lr, func(i, j int) bool {
		return lr[i].y < lr[j].y
	})

	res := 0
	x := -1 << 60
	for _, i := range lr {
		if x+d-1 < i.x {
			res++
			x = i.y
		}
	}
	fmt.Println(res)
}
