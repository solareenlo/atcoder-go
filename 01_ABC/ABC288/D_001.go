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

	var sum [10][200200]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for j := 0; j < k; j++ {
			if i%k == j {
				sum[j][i] = sum[j][i-1] + x
			} else {
				sum[j][i] = sum[j][i-1]
			}
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		v := sum[0][r] - sum[0][l-1]
		f := true
		for i := 0; i < k; i++ {
			if sum[i][r]-sum[i][l-1] != v {
				f = false
				break
			}
		}
		if f {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
