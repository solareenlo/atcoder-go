package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var N int
var A, B [101010]int
var Check [101010]map[int]bool
var G [202020][]int

func IsExist() bool {
	if A[1] != 0 || B[2] != 0 {
		return false
	}
	if A[2] == 0 || B[1] == 0 {
		return false
	}
	if A[2] != B[1] {
		return false
	}
	for i := 1; i <= N; i++ {
		Check[A[i]][B[i]] = true
	}
	for i := 1; i <= N; i++ {
		flag1 := i == 1
		flag2 := i == 2
		for j := -1; j <= 1; j++ {
			if i != 1 && A[i]-1 >= 0 && B[i]+j >= 0 && Check[A[i]-1][B[i]+j] {
				flag1 = true
			}
			if i != 2 && A[i]+j >= 0 && B[i]-1 >= 0 && Check[A[i]+j][B[i]-1] {
				flag2 = true
			}
		}
		if !flag1 || !flag2 {
			return false
		}
	}
	return true
}

func Solve(v []int) int {
	ret := 0
	prv_rem := 0
	for i, j := 0, 0; i < len(v); i = j + 1 {
		for j = i; j < len(v) && B[v[i]] == B[v[j]]; j++ {

		}
		j--
		now := j - i + 1
		other := 0
		if i != 0 && B[v[i]] == B[v[i-1]]+1 {
			other = min(prv_rem, now)
			ret += other
		}
		a := A[v[i]] - 1
		b := B[v[i]] - 1
		if a >= 0 && Check[a][b] {
			ret -= other
			ret += now
			prv_rem = other
		} else {
			prv_rem = now
		}
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := range Check {
		Check[i] = make(map[int]bool)
	}

	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}
	if !IsExist() {
		fmt.Println(-1)
		return
	}
	for i := 1; i <= N; i++ {
		G[A[i]+B[i]] = append(G[A[i]+B[i]], i)
	}
	for i := 1; i <= N+N; i++ {
		sort.Slice(G[i], func(a, b int) bool {
			return B[G[i][a]] < B[G[i][b]]
		})
	}

	ans := 2*N - 2
	for i := 1; i <= N+N; i++ {
		if len(G[i]) != 0 {
			ans -= Solve(G[i])
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
