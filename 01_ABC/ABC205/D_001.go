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
	var n, q int
	fmt.Fscan(in, &n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] -= i
	}

	for i := 0; i < q; i++ {
		var k int
		fmt.Fscan(in, &k)
		m := sort.Search(len(a), func(i int) bool { return k < a[i] })
		fmt.Fprintln(out, m + k)
	}
	out.Flush()
}
