package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, Q, P int
	fmt.Fscan(in, &N, &Q, &P)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, N)
	copy(B, A)
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	B = unique(B)
	for i := 0; i < N; i++ {
		A[i] = lowerBound(B, A[i])
	}
	M := len(B)
	mp := make([]int, M)
	L := make([]int, Q)
	R := make([]int, Q)
	ind := make([]int, Q)
	ans := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &L[i], &R[i])
		L[i]--
	}
	SQ := int(math.Sqrt(float64(Q)))
	now_l := 0
	now_r := 0
	for i := range ind {
		ind[i] = i
	}
	sort.Slice(ind, func(a, b int) bool {
		if L[ind[a]]/SQ != L[ind[b]]/SQ {
			return L[ind[a]] < L[ind[b]]
		}
		if L[ind[a]]/SQ%2 == 1 {
			return R[ind[a]] > R[ind[b]]
		}
		return R[ind[a]] < R[ind[b]]
	})
	p := make([][]int, M)
	for i := 0; i < M; i++ {
		p[i] = append(p[i], 1)
	}
	cur := 0
	add := func(x int) {
		mp[A[x]]++
		if mp[A[x]] >= len(p[A[x]]) {
			c2 := (p[A[x]][len(p[A[x]])-1] * B[A[x]]) % P
			p[A[x]] = append(p[A[x]], c2)
			cur = (cur + p[A[x]][mp[A[x]]]) % P
		} else {
			cur = (cur + p[A[x]][mp[A[x]]]) % P
		}
	}
	del := func(x int) {
		cur = (cur - p[A[x]][mp[A[x]]] + P) % P
		mp[A[x]]--
	}
	for i := 0; i < Q; i++ {
		for now_l > L[ind[i]] {
			now_l--
			add(now_l)
		}
		for now_r < R[ind[i]] {
			add(now_r)
			now_r++
		}
		for now_l < L[ind[i]] {
			del(now_l)
			now_l++
		}
		for now_r > R[ind[i]] {
			now_r--
			del(now_r)
		}
		ans[ind[i]] = cur
	}
	for i := 0; i < Q; i++ {
		fmt.Println(ans[i])
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
