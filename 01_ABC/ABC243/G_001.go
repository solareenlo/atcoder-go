package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Sqrt(x int) int {
	res := int(math.Sqrt(float64(x))) - 1
	for (res+1)*(res+1) <= x {
		res++
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	dp := make([]int, 100100)
	dp[1] = 1
	for i := 2; i <= 100000; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] += dp[j]
		}
	}

	var T int
	fmt.Fscan(in, &T)
	for j := 0; j < T; j++ {
		var n int
		fmt.Fscan(in, &n)
		x := Sqrt(n)
		sum := 0
		for i := 1; (i * i) <= x; i++ {
			sum += (dp[i] * (x - i*i + 1))
		}
		fmt.Fprintln(out, sum)
	}
}
