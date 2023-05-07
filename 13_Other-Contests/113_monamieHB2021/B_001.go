package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	S := make([]int, 0)
	for i := N; i > 0; i /= K {
		S = append(S, i%K)
	}
	S = reverseOrderInt(S)

	L := len(S)
	dp1 := make([]int, L+1)
	dp2 := make([]int, L+1)
	dp2[0] = 1
	for i := 0; i < L; i++ {
		dp1[i+1] = dp1[i]*(K-1) + dp2[i]*S[i]
		if S[i] != K-1 {
			dp2[i+1] = dp2[i]
		} else {
			dp2[i+1] = 0
		}
	}
	fmt.Println(dp1[L])
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
