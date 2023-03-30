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
	sum := a[0]
	res := 0
	if sum == n {
		res++
	}
	l, r := 0, 1
	for l < n {
		if sum >= n && l < n {
			sum -= a[l]
			l++
		} else if sum < n && r < n {
			sum += a[r]
			r++
		} else {
			break
		}
		if sum == n {
			res++
		}
	}
	fmt.Println(res)
}
