package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

type P struct {
	x, y pair
}

var hp, at, de [32]int
var now, NEXT, cl []P
var sumi [32][102][102][102]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &hp[i], &at[i], &de[i])
	}
	now = append(now, P{pair{0, hp[0]}, pair{100, 1}})
	for i := 0; len(now) > 0; i++ {
		for j := 0; j < len(now); j++ {
			a := now[j].x.x
			b := now[j].x.y
			c := now[j].y.x
			d := now[j].y.y
			if a >= n {
				fmt.Println(i)
				return
			}
			if c <= at[a] {
				continue
			}
			ata(5, a, b, c, 1)
			ata(d, a, b, c, d+1)
			aedge(a, min(hp[a], b+de[a]), min(100, c-at[a]+50), 1)
		}
		now = NEXT
		NEXT = cl
	}
	fmt.Println(-1)
}

func ata(x, a, b, c, d int) {
	if x >= b {
		aedge(a+1, hp[a+1], c-at[a], d)
	} else {
		aedge(a, min(hp[a], b-x+de[a]), c-at[a], d)
	}
	return
}

func aedge(a, b, c, d int) {
	if sumi[a][b][c][d] {
		return
	}
	sumi[a][b][c][d] = true
	NEXT = append(NEXT, P{pair{a, b}, pair{c, d}})
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
