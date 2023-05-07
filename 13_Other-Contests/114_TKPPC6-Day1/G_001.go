package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	C := make(map[int][]int)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		C[a] = append(C[a], i)
		if len(C[a]) >= 2 {
			C[a][0], C[a][1] = C[a][1], C[a][0]
		}
	}

	V := make([]int, 0)
	for j := 0; j < m; j++ {
		var b int
		fmt.Fscan(in, &b)
		for _, i := range C[b] {
			k := lowerBound(V, i)
			if k == len(V) {
				V = append(V, i)
			} else {
				V[k] = i
			}
		}
	}
	fmt.Println(len(V))
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
