package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MAX_N = 2000
const MAX_M = 2000

var N, M int
var q float64
var x [MAX_M]int
var p [MAX_M]float64
var logi, logfac [MAX_N + 1]float64
var logspl, logspr [MAX_M + 1]float64
var u [MAX_N + 1][MAX_M + 1]float64

func solve() float64 {
	logfac[0] = 0
	for i := 1; i <= N; i++ {
		logi[i] = math.Log(float64(i))
		logfac[i] = logfac[i-1] + logi[i]
	}
	sum := 0.0
	logspl[0] = -1e8
	logspr[M] = -1e8
	for i := 0; i < M; i++ {
		logspr[i] = math.Log(1.0 - sum)
		sum += 1.0 * p[i] / q
		logspl[i+1] = math.Log(sum)
	}
	for i := 0; i <= N; i++ {
		for j := 0; j <= M; j++ {
			u[i][j] = math.Exp(logspl[j]*float64(i) + logspr[j]*float64(N-i) + logfac[N] - logfac[i] - logfac[N-i])
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j <= M; j++ {
			u[i+1][j] += u[i][j]
		}
	}
	ans := 0.0
	for i := 0; i < N; i++ {
		sum := 0.0
		center := -1
		for j := 0; j < M; j++ {
			if sum < 0.5 {
				center = j
			}
			sum += u[i][j] - u[i][j+1]
		}
		for j := 0; j < M; j++ {
			ans += math.Abs(float64(x[j]-x[center])) * (u[i][j] - u[i][j+1])
		}
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &M, &q)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &x[i], &p[i])
	}
	fmt.Println(solve())
}
