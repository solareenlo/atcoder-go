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

	res := make([]int, 2)
	for i := 0; i < 2; i++ {
		var sign int
		if i == 0 {
			sign = -1
		} else {
			sign = 1
		}
		sum := 0
		for j := 0; j < n; j++ {
			sum += a[j]
			if sign*sum <= 0 {
				res[i] += diff(sum, sign)
				sum = sign
			}
			sign *= -1
		}
	}
	fmt.Println(min(res[0], res[1]))
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
