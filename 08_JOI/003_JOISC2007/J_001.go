package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INFINITY = 1e18

	type pair struct {
		x, y complex128
	}

	var lines [1001]pair

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		var a, b, c, d float64
		fmt.Fscan(in, &a, &b, &c, &d)
		P := complex(a, b)
		Q := complex(c, d)
		lines[i] = pair{P, Q}
	}

	cnt := 1
	VD := make([]float64, 0)
	for i := 0; i < N; i++ {
		bl := false
		P, Q := lines[i].x, lines[i].y
		EP := DivideByReal((Q - P), cmplx.Abs(Q-P))
		for j := 0; j < i; j++ {
			R := (lines[j].x - P) / EP
			S := (lines[j].y - P) / EP

			if same(imag(R), imag(S)) {
				if same(0, imag(R)) {
					bl = true
				}
				VD[j] = INFINITY
				continue
			}

			T := S - R
			a := -imag(R) / imag(T)
			U := complex(a*real(T), 0) + R
			VD[j] = real(U)
		}
		sort.Float64s(VD)
		var x int
		if len(VD) != 0 && VD[0] != INFINITY {
			x = 1
		} else {
			x = 0
		}
		for j := 1; j < i; j++ {
			if VD[j] == INFINITY {
				break
			}
			if !same(VD[j-1], VD[j]) {
				x++
			}
		}
		if !bl {
			cnt += x + 1
		}
		VD = append(VD, 0)
	}
	fmt.Println(cnt)
}

func DivideByReal(c complex128, divisor float64) complex128 {
	return complex(real(c)/divisor, imag(c)/divisor)
}

func same(a, b float64) bool {
	return math.Abs(a-b) < math.Pow(10, -10)
}
