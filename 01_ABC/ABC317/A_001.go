package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, h, x int
	fmt.Fscan(in, &n, &h, &x)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	fmt.Println(lowerBound(p, x-h) + 1)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
