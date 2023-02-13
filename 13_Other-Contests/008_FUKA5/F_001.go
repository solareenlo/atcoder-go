package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type pair struct {
	f float64
	i int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for {
		var n int
		var w float64
		fmt.Fscan(in, &n, &w)
		if n == 0 && w == 0 {
			break
		}
		res1 := 0
		Ev := make([]pair, 0)
		for i := 0; i < n; i++ {
			var x, y float64
			fmt.Fscan(in, &x, &y)
			if x*x+y*y <= w*w {
				res1++
				continue
			}
			h := math.Hypot(x, y)
			Ev = append(Ev, pair{math.Asin(y/h) - math.Asin(w/h), +1})
			Ev = append(Ev, pair{math.Asin(y/h) + math.Asin(w/h), -1})
		}
		sort.Slice(Ev, func(i, j int) bool {
			if Ev[i].f == Ev[j].f {
				return Ev[i].i < Ev[j].i
			}
			return Ev[i].f < Ev[j].f
		})
		res2, cnt := 0, 0
		for _, e := range Ev {
			cnt += e.i
			res2 = max(res2, cnt)
		}
		fmt.Fprintln(out, res1+res2)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
