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
	A := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &C[i])
		A[i]--
	}

	zero := false
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if C[i] != C[j] && C[i] == C[k] {
					zero = true
				}
			}
		}
	}
	if zero {
		fmt.Println(0)
		return
	}
	if C[0] == C[N-1] {
		fmt.Println(-1)
		return
	}

	ans := 0
	var f func(int, int)
	f = func(M, add int) {
		prev := -1
		for i := 0; i < N; i++ {
			curr := A[i] / M
			if i > 0 && ((prev == curr && C[i-1] != C[i]) || (prev != curr && C[i-1] == C[i])) {
				return
			}
			prev = curr
		}
		ans += add
	}

	items := make([]int, 0)
	for i := 1; i <= 100000; i++ {
		items = append(items, i)
	}
	for i := 0; i < N; i++ {
		for col := 0; col < A[i]/100000; col++ {
			items = append(items, A[i]/(col+1)+1)
		}
	}
	sort.Ints(items)
	items = unique(items)
	for i := 0; i+1 < len(items); i++ {
		f(items[i], items[i+1]-items[i])
	}
	fmt.Println(ans)
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
