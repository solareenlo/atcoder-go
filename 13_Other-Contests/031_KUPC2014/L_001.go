package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MOD = 1000000007
const B = 6.0

var ans int
var v []float64
var X float64
var ln [10]float64
var A [4]int = [4]int{2, 3, 5, 7}
var pw4 []int

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := 1; i < 10; i++ {
		ln[i] = math.Log(float64(i))
	}
	var n int
	fmt.Fscan(in, &n)
	pw4 = make([]int, 1011)
	pw4[0] = 1
	for i := 1; i <= n; i++ {
		pw4[i] = pw4[i-1] * 4 % MOD
	}
	a := make([]int, 1011)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	v = make([]float64, 0)
	v = append(v, 1.0)
	v = append(v, 2.0)
	v = append(v, 3.0)
	v = append(v, 4.0)
	v = append(v, 5.0)
	X = 1.0
	for i := n; i >= 1; i-- {
		Go(a[i], i-1)
		if X > 100 {
			X = 1e20
		} else {
			X = math.Min(1e20, math.Pow(float64(a[i]), X))
		}
		if X < B {
			v = append(v, 1.0/X)
		} else {
			ans += pw4[i-1]
		}
	}
	for _, x := range v {
		if x < (1 - (1e-9)) {
			ans += 1.0
		}
	}
	fmt.Println(ans % MOD)
}

func Go(b, Len int) {
	nv := make([]float64, 0)
	for _, k := range v {
		for c := 0; c < 4; c++ {
			nk := (ln[A[c]]/ln[b]*k - 1) * X
			cmp := 0
			if nk > B {
				cmp = 1
			} else if nk < -B {
				cmp = -1
			} else {
				nk = math.Pow(float64(b), nk)
				if nk > B {
					cmp = 1
				} else if nk < 1.0/B {
					cmp = -1
				}
			}
			if cmp == -1 {
				ans += pw4[Len]
			} else if cmp == 0 {
				nv = append(nv, nk)
			}
		}
	}
	v = nv
}
