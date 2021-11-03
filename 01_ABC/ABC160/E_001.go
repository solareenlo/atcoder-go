package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, y, a, b, c int
	fmt.Fscan(in, &x, &y, &a, &b, &c)

	p := make([]int, a)
	for i := range p {
		fmt.Fscan(in, &p[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(p)))

	q := make([]int, b)
	for i := range q {
		fmt.Fscan(in, &q[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(q)))

	r := make([]int, c)
	for i := range r {
		fmt.Fscan(in, &r[i])
	}
	for i := 0; i < x; i++ {
		r = append(r, p[i])
	}
	for i := 0; i < y; i++ {
		r = append(r, q[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))

	res := 0
	for i := 0; i < x+y; i++ {
		res += r[i]
	}

	fmt.Println(res)
}
