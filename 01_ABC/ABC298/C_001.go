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

	const N = 200200

	a := make([][]int, N)
	b := make([][]int, N)

	var n, q int
	fmt.Fscan(in, &n, &q)
	for q > 0 {
		q--
		var op int
		fmt.Fscan(in, &op)
		if op == 1 {
			var v, x int
			fmt.Fscan(in, &v, &x)
			a[x] = append(a[x], v)
			b[v] = append(b[v], x)
		} else if op == 2 {
			var x int
			fmt.Fscan(in, &x)
			sort.Ints(a[x])
			for i := 0; i < len(a[x]); i++ {
				fmt.Fprintln(out, a[x][i])
			}
			fmt.Fprintln(out)
		} else {
			var x int
			fmt.Fscan(in, &x)
			sort.Ints(b[x])
			b[x] = unique(b[x])
			for i := 0; i < len(b[x]); i++ {
				fmt.Fprintln(out, b[x][i])
			}
			fmt.Fprintln(out)
		}
	}
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
