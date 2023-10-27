package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	m := max(N, M)
	A := make([]int, m)
	B := make([]int, m)
	for i := m - N; i < m; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := m - M; i < m; i++ {
		fmt.Fscan(in, &B[i])
	}
	if reflect.DeepEqual(A, B) {
		fmt.Println("Same")
	} else if lessThanSlices(A, B) {
		fmt.Println("X")
	} else {
		fmt.Println("Y")
	}
}

func lessThanSlices(s1, s2 []int) bool {
	minLength := len(s1)
	if len(s2) < minLength {
		minLength = len(s2)
	}

	for i := 0; i < minLength; i++ {
		if s1[i] < s2[i] {
			return true // s1がs2よりも小さい
		} else if s1[i] > s2[i] {
			return false // s1がs2よりも大きい
		}
	}

	if len(s1) == len(s2) {
		return false // スライスは等しい
	} else if len(s1) < len(s2) {
		return true // s1がs2よりも小さい
	}

	return false // s1がs2よりも大きい
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
