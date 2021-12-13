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

	var n int
	fmt.Fscan(in, &n)

	x := make([][2]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		x = append(x, [2]int{a, 1})
		x = append(x, [2]int{a + b, -1})
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i][0] < x[j][0]
	})

	res := make([]int, 200010)
	cnt := 0
	for i := 0; i < len(x)-1; i++ {
		cnt += x[i][1]
		res[cnt] += x[i+1][0] - x[i][0]
	}

	for i := 0; i < n-1; i++ {
		fmt.Fprint(out, res[i+1], " ")
	}
	fmt.Fprintln(out, res[n])
}
