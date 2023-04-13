package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	sort.Ints(v)
	v = reverseOrderInt(v)

	for i := 0; i < n; i = i + 2 {
		x += v[i]
		y += v[i+1]
	}
	if x > y {
		fmt.Println("Takahashi")
	} else if x < y {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Draw")
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
