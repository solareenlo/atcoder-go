package main

import (
	"fmt"
	"math"
)

var P1, P2, P3 float64
var dp [101][101][101][4]float64

func main() {
	var N1, N2, N3 int
	fmt.Scan(&N1, &N2, &N3)
	fmt.Scan(&P1, &P2, &P3)
	P1 /= 100
	P2 /= 100
	P3 /= 100
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}
	fmt.Println(rec(N1, N2, N3, 0))
}

func rec(a, b, c, t int) float64 {
	if a == 0 && b == 0 && c == 0 {
		return 0
	}
	if dp[a][b][c][t] >= 0 {
		return dp[a][b][c][t]
	}
	ret := 1e18
	if t == 0 {
		if a > 0 {
			ret = math.Min(ret, (1.0+rec(a-1, b, c, 1)*P1)/P1)
		}
		if b > 0 {
			ret = math.Min(ret, (1.0+rec(a, b-1, c, 1)*P2)/P2)
		}
		if c > 0 {
			ret = math.Min(ret, (1.0+rec(a, b, c-1, 1)*P3)/P3)
		}
	}
	if t == 1 {
		if a > 0 {
			ret = math.Min(ret, rec(a-1, b, c, 2)*P1+rec(a, b, c, 0)*(1-P1))
		}
		if b > 0 {
			ret = math.Min(ret, rec(a, b-1, c, 2)*P2+rec(a, b, c, 0)*(1-P2))
		}
		if c > 0 {
			ret = math.Min(ret, rec(a, b, c-1, 2)*P3+rec(a, b, c, 0)*(1-P3))
		}
	}
	if t == 2 {
		if a > 0 {
			ret = math.Min(ret, rec(a-1, b, c, 0)*P1+rec(a, b, c, 0)*(1-P1))
		}
		if b > 0 {
			ret = math.Min(ret, rec(a, b-1, c, 3)*P2+rec(a, b, c, 0)*(1-P2))
		}
		if c > 0 {
			ret = math.Min(ret, rec(a, b, c-1, 3)*P3+rec(a, b, c, 0)*(1-P3))
		}
	}
	if t == 3 {
		if a > 0 {
			ret = math.Min(ret, rec(a-1, b, c, 0)*P1+rec(a, b, c, 0)*(1-P1))
		}
		if b > 0 {
			ret = math.Min(ret, rec(a, b-1, c, 0)*P2+rec(a, b, c, 0)*(1-P2))
		}
		if c > 0 {
			ret = math.Min(ret, rec(a, b, c-1, 0)*P3+rec(a, b, c, 0)*(1-P3))
		}
	}
	dp[a][b][c][t] = ret
	return dp[a][b][c][t]
}
