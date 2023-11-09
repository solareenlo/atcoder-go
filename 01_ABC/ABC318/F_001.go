package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var N int
var X, Y, L []int

func check(k int) bool {
	for i := 1; i <= N; i++ {
		Y[i] = abs(X[i] - k)
	}
	sort.Ints(Y[1:])
	for i := 1; i <= N; i++ {
		if Y[i] > L[i] {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var S [80005]int

	X = make([]int, 205)
	L = make([]int, 205)

	fmt.Fscan(in, &N)
	Y = make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &X[i])
	}
	cnt := 0
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &L[i])
		for j := 1; j <= N; j++ {
			cnt++
			S[cnt] = X[j] + L[i]
			cnt++
			S[cnt] = X[j] - L[i] - 1
		}
	}
	sort.Ints(S[1 : cnt+1])
	ans := 0
	for i := 2; i <= cnt; i++ {
		if check(S[i]) {
			ans += S[i] - S[i-1]
		}
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
