package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	R := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &R[i])
	}
	pi := math.Atan(1.0) * 4.0
	l := 0.0
	r := 1000000.0
	for z := 0; z < 100; z++ {
		c := (l + r) / 2.0
		asum := 0.0
		for i := 0; i < N-1; i++ {
			var a float64
			if c < R[i]*R[i+1] {
				a = math.Acos(c / (R[i] * R[i+1]))
			} else {
				a = pi / float64(2*N)
			}
			asum += a
		}

		if asum < pi*2/3 {
			r = c
		} else {
			l = c
		}
	}
	ans := 0.0
	tgt := (l + r) / 2
	for i := 0; i < N-1; i++ {
		var a float64
		if tgt < R[i]*R[i+1] {
			a = math.Acos(tgt / (R[i] * R[i+1]))
		} else {
			a = pi / float64(2*N)
		}
		ans += R[i] * R[i+1] * math.Sin(a) / 2
	}
	ans -= R[0] * R[N-1] * math.Sin(pi*2/3) / 2
	fmt.Println(ans)
}
