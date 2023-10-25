package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var logfac [1 << 17]float64

	var n, m int
	fmt.Fscan(in, &n, &m)
	logfac[0] = 0
	for i := 1; i <= n; i++ {
		logfac[i] = logfac[i-1] + math.Log10(float64(i))
	}
	p := logfac[n] - float64(n)*math.Log10(float64(m))
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		p -= logfac[a]
	}
	fmt.Println(-int(math.Floor(float64(p))))
}
