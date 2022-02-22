package main

import (
	"bufio"
	"fmt"
	"os"
)

func mex(a, b int) int {
	res := 0
	for res == a || res == b {
		res++
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		if i <= 4 {
			a[i] = make([]int, n+1)
		} else {
			a[i] = make([]int, 5)
		}
	}

	t := make([]int, 3)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[1][i])
		t[a[1][i]]++
	}

	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i][1])
		t[a[i][1]]++
	}

	for i := 2; i <= n; i++ {
		for j := 2; j < len(a[i]); j++ {
			a[i][j] = mex(a[i-1][j], a[i][j-1])
			t[a[i][j]]++
		}
	}

	if n > 4 {
		for i := 4; i <= n; i++ {
			t[a[4][i]] += n - i
		}
		for i := 5; i <= n; i++ {
			t[a[i][4]] += n - i
		}
	}

	fmt.Println(t[0], t[1], t[2])
}
