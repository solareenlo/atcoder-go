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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	a = unique(a)

	if len(a) == 1 {
		fmt.Println(0)
	} else {
		cnt := 0
		for i := 1; i < len(a)-1; i++ {
			if (a[i-1] < a[i]) != (a[i] < a[i+1]) {
				cnt++
			}
		}
		fmt.Println(cnt + 2)
	}
}

func unique(a []int) []int {
	if len(a) < 2 {
		return a
	}
	res := make([]int, 0, len(a))
	res = append(res, a[0])
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			res = append(res, a[i])
		}
	}
	return res
}
