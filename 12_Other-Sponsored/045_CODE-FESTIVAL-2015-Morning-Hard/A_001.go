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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	acc := 0
	i := 0
	j := n - 1
	for i != j {
		l := 1 + 2*a[i] + a[i+1]
		r := 1 + 2*a[j] + a[j-1]
		if l < r {
			acc += l
			a[i+2] += 2 + a[i] + a[i+1]
			i += 2
		} else {
			acc += r
			a[j-2] += 2 + a[j] + a[j-1]
			j -= 2
		}
	}
	fmt.Println(acc)
}
