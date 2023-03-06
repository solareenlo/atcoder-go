package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, L, W int
	fmt.Fscan(in, &N, &L, &W)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}
	pos := make([]int, N)
	cnt := make([]int, N)
	for i := 0; i < N; i++ {
		pos[i] = int(1e9)
		cnt[i] = int(1e9)
	}
	pos[0] = 0
	cnt[0] = 0
	j := 0
	for i := 0; i < N-1; i++ {
		if cnt[i] == int(1e9) {
			continue
		}
		for j < N && (i >= j || pos[i]+W > B[j]) {
			j++
		}
		for j < N && B[i]+W >= A[j] {
			pos[j] = max(A[j], pos[i]+W)
			cnt[j] = cnt[i] + 1
			j++
		}
	}
	if cnt[N-1] == 1e9 {
		fmt.Println(-1)
	} else {
		fmt.Println(cnt[N-1])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
