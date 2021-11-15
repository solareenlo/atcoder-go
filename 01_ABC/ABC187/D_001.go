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

	x := make([]int, n)
	sumA := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		sumA += a
		x[i] = 2*a + b
	}
	sort.Sort(sort.Reverse(sort.IntSlice(x)))

	cnt := 0
	for i := 0; i < n; i++ {
		if sumA >= 0 {
			sumA -= x[i]
			cnt++
		}
	}

	fmt.Println(cnt)
}
