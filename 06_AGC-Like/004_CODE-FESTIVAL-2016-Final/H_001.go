package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	n -= 2

	var bg1, bg2 int
	fmt.Fscan(in, &bg1, &bg2)

	a := make([]int, n)
	for i := 1; i <= n-1; i++ {
		fmt.Fscan(in, &a[i])
	}
	sum := 0
	for i := 1; i <= n-1; i++ {
		sum += a[i]
	}
	v := make([]int, sum+1)
	for i := 0; i <= sum; i++ {
		v[sum-i] = i
	}
	for i := 1; i <= n-1; i++ {
		sz := len(v) - 1
		for j := 1; j <= a[i]; j++ {
			v = append(v, v[sz-j])
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		if x >= sum {
			fmt.Fprintln(out, x-sum+bg1-bg2)
		} else {
			fmt.Fprintln(out, v[len(v)-1-x]+bg1-bg2)
		}
	}
}
