package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	S := make([]int, 1<<17)
	for i := 1; i < N; i++ {
		fmt.Fscan(in, &S[i])
	}
	X := make([]int, 1<<17)
	for i := 1; i <= M; i++ {
		fmt.Fscan(in, &X[i])
	}

	a, b, dn := 1, 0, 0
	d := make([]int, 1<<17*10)
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			d[dn] = (X[j] - b) * a
			dn++
		}
		a = -a
		b = S[i] - b
	}
	tmp := d[:dn]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	ans := 0
	var j int
	for i := 0; i < dn; i = j {
		for j = i + 1; j < dn && d[j] == d[i]; j++ {

		}
		ans = max(ans, j-i)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
