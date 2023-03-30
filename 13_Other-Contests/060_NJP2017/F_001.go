package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	N int
	T [100001]int
	X [100001]int
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &T[i], &X[i])
	}

	l, r := 0.0, 2.0*1e6
	for i := 0; i < 45; i++ {
		v := (l + r) / 2.0
		dp := [2][2]float64{}
		idx := 0
		dp[idx] = [2]float64{0.0, 0.0}
		ok := true
		for j := 1; j <= N; j++ {
			idx ^= 1
			vt := v * float64(T[j]-T[j-1])
			l1 := float64(X[j-1]) - vt
			r1 := float64(X[j-1]) + vt
			l2 := dp[idx^1][0] - vt
			r2 := dp[idx^1][1] + vt
			chk1 := (l1 <= float64(X[j])) && (float64(X[j]) <= r1)
			chk2 := (l2 <= float64(X[j])) && (float64(X[j]) <= r2)
			if chk1 && chk2 {
				dp[idx] = [2]float64{math.Min(l1, l2), math.Max(r1, r2)}
			} else if chk1 {
				dp[idx] = [2]float64{l2, r2}
			} else if chk2 {
				dp[idx] = [2]float64{l1, r1}
			} else {
				ok = false
				break
			}
		}
		if ok {
			r = v
		} else {
			l = v
		}
	}
	fmt.Printf("%.6f\n", r)
}
