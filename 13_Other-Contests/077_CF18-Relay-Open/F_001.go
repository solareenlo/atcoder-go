package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)
	ans := 0.0
	var t [100000]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &t[i])
		ans += float64(t[i])
	}
	var w [300]float64
	for i := 0; i < K; i++ {
		for j := i; j < K; j++ {
			w[i] += float64(j-i) * (1.0 / float64(K))
		}
		for j := 0; j < K; j++ {
			w[i] += float64(K-i+j) * (float64(i) / float64(K)) * (1.0 / float64(K))
		}
	}
	var prob, tmp [300]float64
	prob[0] = 1
	for i := 1; i <= N; i++ {
		for j := 0; j < K; j++ {
			ans += prob[j] * w[j]
		}
		sum := 0.0
		add := 0.0
		for j := 0; j < K; j++ {
			add += prob[j] * float64(j) / float64(K) / float64(K)
		}
		for j := 0; j < K; j++ {
			sum += prob[j]
			tmp[j] = sum/float64(K) + add
		}
		for j := 0; j < K; j++ {
			prob[(j+t[i-1])%K] = tmp[j]
			tmp[j] = 0
		}
	}
	fmt.Println(ans)
}
