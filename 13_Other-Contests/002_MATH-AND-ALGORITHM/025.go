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

	a := make([]float64, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	b := make([]float64, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	res := 0.0
	for i := 0; i < n; i++ {
		res += a[i]*(1.0/3.0) + b[i]*(2.0/3.0)
	}
	fmt.Println(res)
}
