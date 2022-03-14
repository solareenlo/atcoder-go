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
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[i] = pair{x + y, x - y}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x
	})

	cur := -1 << 60
	ans := 0
	for i := 1; i <= n; i++ {
		if cur <= a[i].y {
			ans++
			cur = a[i].x
		}
	}
	fmt.Println(ans)
}
