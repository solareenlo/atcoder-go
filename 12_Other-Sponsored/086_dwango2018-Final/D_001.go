package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 62
const MOD = 1000000007

type Matrix struct {
	f [N][N]int
}

var A, S [N]Matrix

func (mat *Matrix) mul(a Matrix) Matrix {
	var res Matrix
	for j := 0; j < 60; j++ {
		for i := 0; i < 60; i++ {
			for k := 0; k < 60; k++ {
				res.f[i][k] = (res.f[i][k] + mat.f[i][j]*a.f[j][k]%MOD) % MOD
			}
		}
	}
	return res
}

var ans [N]int

func mul(a Matrix) {
	var tmp [N]int
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			tmp[i] = (tmp[i] + ans[j]*a.f[j][i]) % MOD
		}
	}
	for i := 0; i < 60; i++ {
		ans[i] = tmp[i]
	}
}

func Init(x int) {
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			A[i].f[j][j] = 1
			if j > i {
				continue
			}
			for k := j + 1; k < 60; k++ {
				if ((x >> (k - j - 1)) & 1) != 0 {
					A[i].f[k][j] = 1
				}
			}
		}
	}
	S[1] = A[0]
	for i := 2; i < 60; i++ {
		tmp := S[i-1].mul(A[i-1])
		S[i] = tmp.mul(S[i-1])
	}
	for i := 1; i < 60; i++ {
		S[i] = S[i].mul(A[i])
	}
	S[0] = A[0]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Init(1)

	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var n int
		fmt.Fscan(in, &n)
		for i := 0; i < 60; i++ {
			ans[i] = 1
		}
		n--
		for j := 60 - 1; j >= 0; j-- {
			if (n>>j)&1 != 0 {
				mul(S[j])
			}
		}
		fmt.Println(ans[0])
	}
}
