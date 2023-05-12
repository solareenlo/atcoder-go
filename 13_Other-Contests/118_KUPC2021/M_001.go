package main

import "fmt"

const MOD = 998244353

func main() {
	var A Matrix
	for i := 4; i < 8; i++ {
		A[i-4][i] += 1
	}
	for d := 1; d <= 9; d++ {
		A[0][0] += 1
		A[1][1] += 1
		A[1][2] += 9
		A[1][3] += d
		A[2][2] += 10
		A[2][3] += d
		A[3][3] += 1
		A[4][0] += 1
		A[5][0] += d
		A[5][1] += 1
		A[6][0] += d
		A[7][0] += 1
		A[4][0] += 1
		A[5][1] += 1
		A[5][2] += d - 1
		A[6][2] += d
		A[7][2] += 1
	}
	var N int
	fmt.Scan(&N)
	A = A.pow(N - 1)
	fmt.Println((A[1][0]*9%MOD + A[1][1]*45%MOD + A[1][2]*45%MOD + A[1][3]*9%MOD) % MOD)
}

type Matrix [8][8]int

func eye() Matrix {
	var res Matrix
	for i := 0; i < 8; i++ {
		res[i][i] = 1
	}
	return res
}

func (dat Matrix) mul(A Matrix) Matrix {
	var res Matrix
	for i := 0; i < 8; i++ {
		for k := 0; k < 8; k++ {
			for j := 0; j < 8; j++ {
				res[i][j] = (res[i][j] + dat[i][k]*A[k][j]%MOD) % MOD
			}
		}
	}
	return res
}

func (a Matrix) pow(n int) Matrix {
	res := eye()
	for n > 0 {
		if (n & 1) != 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
		n >>= 1
	}
	return res
}
