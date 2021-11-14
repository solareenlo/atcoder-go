package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(in, &n, &w)

	N := make([]int, 200005)
	s := make([]int, 200005+1)
	for i := 0; i < n; i++ {
		var s, t, p int
		fmt.Fscan(in, &s, &t, &p)
		N[s] += p
		N[t] -= p
	}

	for i := 0; i < 200005; i++ {
		s[i+1] = s[i] + N[i]
	}

	if w >= max(s...) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
