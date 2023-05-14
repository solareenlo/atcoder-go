package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	var p, q [110]float64
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
		p[i] /= 100.0
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &q[i])
		q[i] /= 100.0
	}
	ans := 0.0
	for i := k - 1; i >= 0; i-- {
		fl := 0
		ma := 0.0
		for j := 0; j < n; j++ {
			if p[j]*q[j] > ma {
				ma = p[j] * q[j]
				fl = j
			}
		}
		ans += ma * float64(i)
		p[fl] *= (1 - q[fl])
	}
	fmt.Println(ans)
}
