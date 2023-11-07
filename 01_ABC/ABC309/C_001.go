package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	A := make([]pair, n)
	c := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i].x, &A[i].y)
		c += A[i].y
	}
	if c <= k {
		fmt.Println(1)
		return
	}
	sortPair(A)
	for _, tmp := range A {
		a, b := tmp.x, tmp.y
		c -= b
		if c <= k {
			fmt.Println(a + 1)
			return
		}
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
