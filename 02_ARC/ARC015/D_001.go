package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t, n int
	var p float64
	fmt.Fscan(in, &t, &n, &p)

	e := [100100]float64{}
	for i := 0; i < n; i++ {
		var q float64
		var x, t int
		fmt.Fscan(in, &q, &x, &t)
		e[t-1] += float64(x-1) * q * p
	}

	e[100100-1] = 1.0
	for i := 100100 - 2; i >= 0; i-- {
		e[i] += e[i+1]
	}

	ans := 1.0
	cur := 1.0
	for i := 0; i < t-1; i++ {
		cur *= e[i]
		ans += cur
	}
	fmt.Println(ans)
}
