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
	A := make([]pair, 0)
	B := make([]int, 1441)
	for i := 0; i < n; i++ {
		var t, d int
		fmt.Fscan(in, &t, &d)
		A = append(A, pair{d, t})
	}
	sortPair(A)
	for _, tmp := range A {
		d := tmp.x
		t := tmp.y
		for i := d; i >= t; i-- {
			B[i] = max(B[i], B[i-t]+1)
		}
	}
	ans := 0
	for i := range B {
		ans = max(ans, B[i])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
