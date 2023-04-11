package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f [3010]float64
	f[0] = 1.0
	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var k float64
		fmt.Fscan(in, &k)
		for j := i - 1; j >= 0; j-- {
			f[j+1] += f[j] * k
			f[j] *= (1 - k)
		}
	}
	s := 0.0
	for i := n; i > n/2; i-- {
		s += f[i]
	}
	fmt.Println(s)
}
