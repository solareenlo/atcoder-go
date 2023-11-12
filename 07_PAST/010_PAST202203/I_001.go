package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([]pair, N)
	B := make([]pair, N)
	C := make([]pair, N)
	D := make([]pair, N)
	mnx := 1145141919
	mny := 1145141919
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i].x, &A[i].y)
		mnx = min(A[i].x, mnx)
		mny = min(A[i].y, mny)
	}
	sortPair(A)
	mxx := -1145141919
	mxy := -1145141919
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &B[i].x, &B[i].y)
		mxx = max(B[i].x, mxx)
		mxy = max(B[i].y, mxy)
	}
	sortPair(B)
	for i := 0; i < N; i++ {
		C[i] = pair{mnx + mxx - A[i].x, A[i].y}
	}
	sortPair(C)
	for i := 0; i < N; i++ {
		D[i] = pair{A[i].x, mny + mxy - A[i].y}
	}
	sortPair(D)
	if equalSlices(A, B) || equalSlices(B, C) || equalSlices(B, D) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func equalSlices(s1, s2 []pair) bool {
	minLength := len(s1)
	if len(s2) < minLength {
		minLength = len(s2)
	}

	for i := 0; i < minLength; i++ {
		if lessThanPair(s1[i], s2[i]) {
			return false // s1がs2よりも小さい
		} else if lessThanPair(s2[i], s1[i]) {
			return false // s1がs2よりも大きい
		}
	}

	if len(s1) == len(s2) {
		return true // スライスは等しい
	} else if len(s1) < len(s2) {
		return false // s1がs2よりも小さい
	}

	return false // s1がs2よりも大きい
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		return lessThanPair(tmp[i], tmp[j])
	})
}

func lessThanPair(l, r pair) bool {
	if l.x == r.x {
		return l.y < r.y
	}
	return l.x < r.x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
