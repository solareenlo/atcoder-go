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

	const N = 200020

	var n, m int
	fmt.Fscan(in, &n, &m)
	v := make([][]int, N)
	for m > 0 {
		m--
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", len(v[i]))
		sort.Ints(v[i])
		for _, x := range v[i] {
			fmt.Fprintf(out, "%d ", x)
		}
		fmt.Fprintln(out)
	}
}
