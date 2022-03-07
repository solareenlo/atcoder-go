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

	p := make([]int, n)
	for i := range p {
		fmt.Fscan(in, &p[i])
	}
	q := make([]int, n+1)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		q[a] = i
	}

	r := make([]int, 0)
	for i := 0; i < n; i++ {
		v := make([]int, 0)
		for j := p[i]; j <= n; j += p[i] {
			v = append(v, q[j])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		for _, x := range v {
			r = append(r, x)
		}
	}

	s := make([]int, 0)
	for _, p := range r {
		idx := lowerBound(s, p)
		if len(s) == 0 || len(s) == idx {
			s = append(s, p)
		} else {
			s[idx] = p
		}
	}
	fmt.Println(len(s))
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
