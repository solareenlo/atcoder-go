package main

import (
	"fmt"
)

var used [2020]bool
var dist [2020][2020]int
var X [2020][]int

func dfs(pos, fro, depth int) {
	if used[pos] == true {
		return
	}
	used[pos] = true
	dist[fro][pos] = depth
	for i := 0; i < len(X[pos]); i++ {
		dfs(X[pos][i], fro, depth+1)
	}
}

func main() {

	type pair struct {
		x, y int
	}

	var N, Q int
	fmt.Scan(&N, &Q)
	var A, B [2020]int
	for i := 1; i <= N-1; i++ {
		fmt.Scan(&A[i], &B[i])
		X[A[i]] = append(X[A[i]], B[i])
		X[B[i]] = append(X[B[i]], A[i])
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			used[j] = false
		}
		dfs(i, i, 0)
	}
	S := make([]int, 0)
	for i := 1; i <= N; i++ {
		S = append(S, i)
	}
	for len(S) >= 2 {
		maxn := 1 << 30
		maxid := pair{-1, -1}
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if dist[i][j] > 2 {
					continue
				}
				s1, s2, s3 := 0, 0, 0
				for _, k := range S {
					if dist[i][k] < dist[j][k] {
						s1++
					}
					if dist[i][k] > dist[j][k] {
						s2++
					}
					if dist[i][k] == dist[j][k] {
						s3++
					}
				}
				if maxn > max(s1, s2, s3) {
					maxn = max(s1, s2, s3)
					maxid = pair{i, j}
				}
			}
		}
		fmt.Println("?", maxid.x, maxid.y)
		var V int
		a1 := maxid.x
		a2 := maxid.y
		fmt.Scan(&V)
		T := make([]int, 0)
		if V == maxid.x {
			for _, i := range S {
				if dist[a1][i] < dist[a2][i] {
					T = append(T, i)
				}
			}
		} else if V == maxid.y {
			for _, i := range S {
				if dist[a1][i] > dist[a2][i] {
					T = append(T, i)
				}
			}
		} else {
			for _, i := range S {
				if dist[a1][i] == dist[a2][i] {
					T = append(T, i)
				}
			}
		}
		S = T
	}
	fmt.Println("!", S[0])
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
