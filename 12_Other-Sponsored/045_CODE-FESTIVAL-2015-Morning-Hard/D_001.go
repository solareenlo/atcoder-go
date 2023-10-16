package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)
	ant := make([]pair, 0)
	for i := 0; i < N; i++ {
		var x, s float64
		var d string
		fmt.Fscan(in, &x, &s, &d)
		if d == "L" {
			s = -s
		}
		ant = append(ant, pair{x, s})
	}
	sortPair(ant)

	l := 0.0
	r := 1e10
	for t := 0; t < 100; t++ {
		m := (l + r) / 2.0
		Len := make([]float64, 0)
		for _, p := range ant {
			pos := p.x + m*p.y
			i := lowerBound(Len, pos)
			if i == len(Len) {
				Len = append(Len, pos)
			} else {
				Len[i] = pos
			}
		}
		if len(Len) >= N-K {
			l = m
		} else {
			r = m
		}
	}

	if r == 1e10 {
		fmt.Println("Infinity")
	} else {
		fmt.Printf("%.12f\n", r)
	}
}

type pair struct {
	x, y float64
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func lowerBound(a []float64, x float64) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
