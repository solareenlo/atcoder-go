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

	var v [200000][]int
	var pp [800000]pair

	var h, w int
	fmt.Fscan(in, &h, &w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var a int
			fmt.Fscan(in, &a)
			v[i] = append(v[i], a)
		}
	}
	p := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j < w-1 {
				x := v[i][j]
				y := v[i][j+1]
				if x > y {
					x, y = y, x
				}
				pp[p] = pair{x, y}
				p++
			}
			if i < h-1 {
				x := v[i][j]
				y := v[i+1][j]
				if x > y {
					x, y = y, x
				}
				pp[p] = pair{x, y}
				p++
			}
		}
	}
	sortPair(pp[:p])
	for i := 0; i < p; i++ {
		fmt.Fprintf(out, "%d %d\n", pp[i].x, pp[i].y)
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
