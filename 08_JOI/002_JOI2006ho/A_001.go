package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	place := make([]pair, m)
	for i := 0; i < m; i++ {
		place[i].y = i + 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var res int
			fmt.Fscan(in, &res)
			place[j].x -= res
		}
	}
	sortPair(place)

	for i := 0; i < m; i++ {
		fmt.Fprint(out, place[i].y)
		if i < m-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
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
