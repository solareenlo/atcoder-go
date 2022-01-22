package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, L int
	fmt.Fscan(in, &n, &L)

	x := make([]int, n)
	d := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &d[i])
	}

	index := 0
	ans := 0
	for index < n {
		l, r := 0, 0
		jumpl, jumpr := 0, 0
		left := 0
		for index < n && d[index] == "R" {
			jumpl += (x[index] - left - 1) * l
			left = x[index]
			l++
			index++
		}
		if index == n {
			ans += jumpl + (L-left)*l
			break
		}
		right := x[index]
		tail := x[index]
		index++
		r++
		for index < n && d[index] == "L" {
			jumpr += x[index] - tail - 1
			tail++
			r++
			index++
		}
		if l > r {
			ans += jumpl + jumpr + (right-left-1)*l
		} else {
			ans += jumpl + jumpr + (right-left-1)*r
		}
	}

	fmt.Println(ans)
}
