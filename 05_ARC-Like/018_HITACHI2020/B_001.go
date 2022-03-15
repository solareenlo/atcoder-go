package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, m int
	fmt.Fscan(in, &a, &b, &m)

	A := make([]int, a+1)
	B := make([]int, b+1)
	ma := 1 << 60
	for i := 1; i <= a; i++ {
		fmt.Fscan(in, &A[i])
		ma = min(ma, A[i])
	}

	mb := 1 << 60
	for i := 1; i <= b; i++ {
		fmt.Fscan(in, &B[i])
		mb = min(mb, B[i])
	}

	ans := mb + ma
	for i := 0; i < m; i++ {
		var l, r, c int
		fmt.Fscan(in, &l, &r, &c)
		ans = min(ans, A[l]+B[r]-c)
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
