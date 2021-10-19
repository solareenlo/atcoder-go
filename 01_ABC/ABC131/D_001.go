package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct{ b, a int }

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	ba := make([]pair, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		ba[i] = pair{b, a}
	}
	sort.Slice(ba, func(i, j int) bool {
		return ba[i].b < ba[j].b
	})

	sum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			sum[i] = ba[i].a
		} else {
			sum[i] = sum[i-1] + ba[i].a
		}
	}

	for i := 0; i < n; i++ {
		if ba[i].b < sum[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
