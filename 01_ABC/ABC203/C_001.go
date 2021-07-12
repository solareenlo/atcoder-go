package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	f, s int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]pair, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[i] = pair{x, y}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].f < a[j].f
	})

	for _, x := range a {
		if k < x.f {
			break
		}
		k += x.s
	}
	fmt.Println(k)
}
