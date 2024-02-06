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
	x := make([]pair, n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Fscan(in, &s)
		for i := 0; i < n; i++ {
			if s[i] == 'o' {
				x[i].x++
			}
		}
		x[j].y = j + 1
	}
	sortPair(x)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", x[i].y)
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
