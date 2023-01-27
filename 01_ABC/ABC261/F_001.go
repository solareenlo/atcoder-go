package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var tr []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	const N = 300005
	a := make([]int, N)
	v := make([][]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		v[a[i]] = append(v[a[i]], i)
	}
	b := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	tr = make([]int, N)
	ans := 0
	for i := 1; i <= n; i++ {
		l := len(v[i])
		for j := 0; j < l; j++ {
			ans -= sum(b[v[i][j]] + 1)
			add(b[v[i][j]], 1)
		}
		for j := 0; j < l; j++ {
			add(b[v[i][j]], -1)
		}
	}
	for i := 1; i <= n; i++ {
		ans += sum(b[i] + 1)
		add(b[i], 1)
	}

	fmt.Println(ans)
}

func lowbit(x int) int {
	return x & (-x)
}
func add(x, y int) {
	for ; x > 0; x -= lowbit(x) {
		tr[x] += y
	}
}

func sum(x int) int {
	s := 0
	for ; x <= n; x += lowbit(x) {
		s += tr[x]
	}
	return s
}
