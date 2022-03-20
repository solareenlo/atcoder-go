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

	sumB := 0.0
	for i := 0; i < n; i++ {
		var b float64
		fmt.Fscan(in, &b)
		sumB += b
	}

	sumR := 0.0
	for i := 0; i < n; i++ {
		var r float64
		fmt.Fscan(in, &r)
		sumR += r
	}

	fmt.Println(sumB/float64(n) + sumR/float64(n))
}
