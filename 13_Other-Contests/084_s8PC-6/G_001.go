package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, P int
	fmt.Fscan(in, &N, &P)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	ans := Solve(N, P, A)
	fmt.Println(ans)
}

func Solve(N, P int, A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	if P == 1 {
		ans := 1
		for i := 0; i < 3; i++ {
			limit := A[i] + 1
			for j := i + 1; j < N; j++ {
				if A[j]+N >= limit {
					ans += comb(j-i-1, 2-i)
				}
			}
		}
		return ans
	}
	if P == 2 {
		ans := 0
		for i := 0; i <= 3; i++ {
			limit := A[i] + 1
			ptr := 1
			tmp0 := 2
			if i <= 1 {
				tmp0 = N
			}
			for j := 1; j < tmp0; j++ {
				if j == i {
					continue
				}
				for ptr != N && A[ptr]+N >= A[j]+2 && A[ptr]+N >= limit {
					ptr++
				}
				if i == 3 && ptr > 3 {
					ptr = 3
				}
				if A[j]+N < limit {
					break
				}
				tmp1 := 1
				if i == 0 {
					tmp1 = j - 1
				}
				tmp2 := 0
				if j+1 <= i && i < ptr {
					tmp2 = 1
				}
				ans += tmp1 * (ptr - (j + 1) - tmp2)
				if ptr != N && ptr != i && A[ptr]+N == A[j]+1 && A[j]+1 >= A[i]+2 {
					ans += tmp1
				}
			}
		}
		return ans
	}
	return -int(1e18)
}

func comb(a, b int) int {
	ans := 1
	for i := 1; i <= b; i++ {
		ans *= a - i + 1
		ans /= i
	}
	return ans
}
