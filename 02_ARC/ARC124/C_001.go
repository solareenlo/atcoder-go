package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	mpl = make([]map[pair]int, 57)
	a   = make([]int, 57)
	b   = make([]int, 57)
)

type pair struct{ x, y int }

func dfs(num, i, j int) int {
	if num == n {
		return i / gcd(i, j) * j
	}
	if _, ok := mpl[num][pair{i, j}]; ok {
		return mpl[num][pair{i, j}]
	}
	mpl[num][pair{i, j}] =
		max(dfs(num+1, gcd(i, a[num]), gcd(j, b[num])),
			dfs(num+1, gcd(i, b[num]), gcd(j, a[num])))
	return mpl[num][pair{i, j}]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 0; i != n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	for i := range mpl {
		mpl[i] = make(map[pair]int, 1)
	}

	fmt.Println(dfs(0, 0, 0))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
