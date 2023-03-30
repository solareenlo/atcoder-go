package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	r, sum, res := 0, 0, 0
	for l := 0; l < n; l++ {
		for r < n && sum < n {
			sum += a[r]
			r++
		}
		if sum == n {
			res++
		}
		if l == r {
			r++
		} else {
			sum -= a[l]
		}
	}
	fmt.Println(res)
}
