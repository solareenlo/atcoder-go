package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

var x, y, s, t [1010]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var M int
	fmt.Fscan(in, &M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		s[i] = 0
		t[i] = 0
	}
	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[M+i], &y[M+i], &s[M+i], &t[M+i])
	}
	A := make([][]pair, 1010)
	if M >= 1 && ok(M-1, M) {
		A[M] = append(A[M], minmax(M-2, M-3))
	}
	if M >= 2 && ok(M-2, M) {
		A[M] = append(A[M], minmax(M-1, M-3))
	}
	if M >= 3 && ok(M-3, M) {
		A[M] = append(A[M], minmax(M-1, M-2))
	}
	for i := M; i+1 < M+N; i++ {
		sort.Slice(A[i], func(a, b int) bool {
			if A[i][a].x == A[i][b].x {
				return A[i][a].y < A[i][b].y
			}
			return A[i][a].x < A[i][b].x
		})
		A[i] = unique(A[i])
		for _, p := range A[i] {
			if x[i] == x[i+1] && y[i] == y[i+1] {
				A[i+1] = append(A[i+1], p)
			} else {
				if p.x >= 0 && ok(p.x, i+1) {
					A[i+1] = append(A[i+1], minmax(i, p.y))
				}
				if p.y >= 0 && ok(p.y, i+1) {
					A[i+1] = append(A[i+1], minmax(i, p.x))
				}
			}
		}
	}
	if len(A[N+M-1]) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func minmax(a, b int) pair {
	if a > b {
		a, b = b, a
	}
	return pair{a, b}
}

func unique(a []pair) []pair {
	occurred := map[pair]bool{}
	result := []pair{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].x == result[j].x {
			return result[i].y < result[j].y
		}
		return result[i].x < result[j].x
	})
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func ok(i, j int) bool {
	return (x[i]-x[j])*(x[i]-x[j])+(y[i]-y[j])*(y[i]-y[j]) <= (s[j]-t[i])*(s[j]-t[i])
}
