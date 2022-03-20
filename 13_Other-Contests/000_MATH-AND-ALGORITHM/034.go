package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	ans := 1e+7

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = math.Min(ans, math.Pow(x[i]-x[j], 2)+math.Pow(y[i]-y[j], 2))
		}
	}

	fmt.Println(math.Sqrt(ans))
}
