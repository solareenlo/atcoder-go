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
	A := make([]int, 0)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		c := lowerBound(A, a)
		if c == len(A) {
			A = append(A, a)
		} else {
			A[c] = a
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
