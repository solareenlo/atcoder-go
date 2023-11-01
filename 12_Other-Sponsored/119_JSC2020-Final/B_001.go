package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	ssort := N - 1
	for ssort > 0 && A[ssort-1] <= A[ssort] {
		ssort--
	}

	A = append(A, A...)
	Max := 0
	for i := range A {
		Max = max(Max, A[i])
	}

	s := make([]int, 1)
	s[0] = 2 * N
	ans, sum := int(1e18), 0
	for i := 2*N - 1; i >= 0; i-- {
		for len(s) > 1 && A[i] > A[s[len(s)-1]] {
			k := s[len(s)-1]
			s = s[:len(s)-1]
			sum -= A[k] * (s[len(s)-1] - k)
		}
		sum += A[i] * (s[len(s)-1] - i)
		s = append(s, i)
		if i < N {
			res := sum - (N-i)*Max + i
			if i < ssort {
				res += N
			}
			ans = min(ans, res)
		}
	}
	sum = accumulate(A[:N])
	ans -= sum
	fmt.Println(ans)
}

func accumulate(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
