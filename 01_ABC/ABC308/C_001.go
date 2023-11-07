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

	type node struct {
		A, B, o int
	}

	var N int
	fmt.Fscan(in, &N)
	a := make([]node, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i].A, &a[i].B)
		a[i].B += a[i].A
		a[i].o = i
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].A*a[j].B != a[j].A*a[i].B {
			return a[i].A*a[j].B > a[j].A*a[i].B
		}
		return a[i].o < a[j].o
	})
	for i := 0; i < N; i++ {
		fmt.Fprintf(out, "%d ", a[i].o+1)
	}
}
