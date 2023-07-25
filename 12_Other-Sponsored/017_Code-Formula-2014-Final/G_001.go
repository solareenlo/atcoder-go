package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var Q [3][]int
var ret []pair

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		var r int
		fmt.Fscan(in, &r)
		Q[0] = append(Q[0], r-1)
	}

	SortTo(0, 1, 0, N)

	fmt.Println(len(ret))
	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i].x, ret[i].y)
	}
}

func SortTo(s, d, L, R int) {
	if R-L == 1 {
		if s != d {
			Move(s, d)
		}
		return
	}

	var p, q int
	if s != d {
		p = 3 - s - d
		q = d
	} else {
		p = (s + 1) % 3
		q = (s + 2) % 3
	}

	mid := (L + R) / 2
	for i := 0; i < R-L; i++ {
		v := Top(s)
		if v < mid {
			Move(s, p)
		} else {
			Move(s, q)
		}
	}

	SortTo(q, d, mid, R)
	SortTo(p, d, L, mid)
}

func Move(s, d int) {
	ret = append(ret, pair{s + 1, d + 1})
	Q[d] = append(Q[d], Top(s))
	Q[s] = Q[s][:len(Q[s])-1]
}

func Top(s int) int { return Q[s][len(Q[s])-1] }
