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
	B := make([]pair, 0)
	A := make([]int, 0)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		B = append(B, pair{x, -y})
	}
	sortPair(B)
	for _, tmp := range B {
		y := tmp.y
		c := lowerBound(A, -y)
		if c == len(A) {
			A = append(A, -y)
		} else {
			A[c] = -y
		}
	}
	fmt.Println(len(A))
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
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
