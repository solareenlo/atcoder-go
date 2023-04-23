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
	a := make([]pair, n)
	for i := 0; i < n; i++ {
		var x int
		var c string
		fmt.Fscan(in, &x, &c)
		a[i] = pair{-int(c[0]), x}
	}
	sortPair(a)

	for i := 0; i < n; i++ {
		fmt.Println(a[i].y)
	}
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
