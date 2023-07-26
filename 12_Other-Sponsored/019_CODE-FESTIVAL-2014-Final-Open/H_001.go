package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	var s string
	fmt.Scan(&s)
	num := make([]int, N+1)
	mem := make([]int, N)
	num[0] = K
	minNum := 0
	maxNum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			mem[i] = minNum + 1
			num[minNum]--
			num[minNum+1]++
			maxNum = max(maxNum, minNum+1)
		} else {
			mem[i] = maxNum + 1
			num[maxNum]--
			num[maxNum+1]++
			maxNum++
		}
		if num[minNum] == 0 {
			minNum++
		}
	}
	addProb := make([]float64, N+1)
	res := make([]float64, N)
	for i := N - 1; i >= 0; i-- {
		m := mem[i]
		res[i] = float64(m) + addProb[m]
		num[m]--
		addProb[m-1] *= float64(num[m-1])
		num[m-1]++
		addProb[m-1] += addProb[m] + 1
		addProb[m-1] /= float64(num[m-1])
	}
	for i := 0; i < N; i++ {
		fmt.Println(res[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
