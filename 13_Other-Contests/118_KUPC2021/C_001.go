package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([]pair, 0)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		A = append(A, pair{a, -1})
	}
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		A = append(A, pair{a, 1})
	}
	sortPair(A)

	ans := int(9e18)
	cum := A[len(A)-1].x * 2
	c := 0
	prv := 0
	for _, p := range A {
		if c >= 0 {
			cum += -(p.x - prv)
		} else {
			cum += p.x - prv
		}
		prv = p.x
		c += p.y
		ans = min(ans, cum)
	}
	fmt.Println(ans)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
