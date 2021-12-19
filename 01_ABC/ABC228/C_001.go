package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	k -= 1
	p := make([]int, n)
	for i := range p {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		p[i] = a + b + c
	}

	q := make([]int, n)
	copy(q, p)
	sort.Sort(sort.Reverse(sort.IntSlice(q)))

	for _, x := range p {
		if x+300 >= q[k] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
