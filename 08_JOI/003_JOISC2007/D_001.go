package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 10005

	var N int
	fmt.Fscan(in, &N)
	p := make([]int, N+1)
	for i := range p {
		p[i] = MX
	}
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		b := lowerBound(p, a)
		p[b] = a
	}
	fmt.Println(lowerBound(p, MX))
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
