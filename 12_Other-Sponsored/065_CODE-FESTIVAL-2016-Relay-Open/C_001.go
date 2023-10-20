package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	n = pow(2, n)
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		q = append(q, x)
	}
	for len(q) > 1 {
		x := q[0]
		q = q[1:]
		y := q[0]
		q = q[1:]
		if x > y {
			q = append(q, x-y)
		} else if x == y {
			q = append(q, x)
		} else {
			q = append(q, y-x)
		}
	}
	fmt.Println(q[0])
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
