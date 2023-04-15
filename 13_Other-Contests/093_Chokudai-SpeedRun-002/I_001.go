package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	i0 := 0
	for i := 1; i < n; i++ {
		if a[i]*b[i] > a[i0]*b[i0] {
			i0 = i
		}
	}

	ok := true
	for j := 0; j < n; j++ {
		if j != i0 && (a[i0]+b[j]-1)/b[j] <= (a[j]+b[i0]-1)/b[i0] {
			ok = false
		}
	}
	if ok {
		fmt.Println(i0 + 1)
	} else {
		fmt.Println(-1)
	}
}
