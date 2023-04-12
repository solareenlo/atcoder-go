package main

import (
	"bufio"
	"fmt"
	"os"
)

var R [22]int
var E [3000009][10]float64
var FinalAns [10]float64
var N, D, K int
var cnt [10][10]float64
var S string

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &D, &K, &S)
	for i := 0; i < 22; i++ {
		R[i] = i % int(D)
	}
	calc()
	for i := 0; i < N; i += D {
		for j := i; j < i+D; j++ {
			cnt[S[j]-'0'][R[(j%D)-int(S[j]-'0')+D]] += float64(j + 1)
		}
		for j := 0; j < D; j++ {
			for k := 0; k < D; k++ {
				FinalAns[j] += 1.0 * cnt[j][k] * E[i/D+1][k]
			}
		}
	}
	for i := 0; i < D; i++ {
		fmt.Println(FinalAns[i])
	}
}

func calc() {
	vec := make([]int, 1)
	vec[0] = 1
	for {
		cx := vec[len(vec)-1]
		cx = max(cx+1, cx*101/100)
		if cx > N/D {
			break
		}
		vec = append(vec, cx)
	}
	vec = append(vec, N/D)
	for i := 0; i < len(vec); i++ {
		G := solve(D, K, 1.0/(float64(D)*float64(vec[i])))
		for j := 0; j < D; j++ {
			E[vec[i]][j] = G[j]
		}
	}
	for i := 0; i < len(vec)-1; i++ {
		cl := vec[i]
		cr := vec[i+1]
		for j := cl + 1; j <= cr-1; j++ {
			for k := 0; k < D; k++ {
				E[j][k] = float64(cr-j)*E[cl][k] + float64(j-cl)*E[cr][k]
				E[j][k] /= float64(cr - cl)
			}
		}
	}
}

func solve(D, K int, P float64) []float64 {
	F := make([]float64, D)
	G := make([]float64, D)
	G[0] = 1.0
	F[0] = 1.0 - 2.0*P
	F[1] = P
	F[D-1] = P
	for i := 0; i < 30; i++ {
		if (K & (1 << i)) != 0 {
			G = mul(G, F)
		}
		F = mul(F, F)
	}
	return G
}

func mul(V1, V2 []float64) []float64 {
	V3 := make([]float64, len(V1))
	for i := 0; i < len(V1); i++ {
		for j := 0; j < len(V2); j++ {
			V3[R[i+j]] += V1[i] * V2[j]
		}
	}
	return V3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
