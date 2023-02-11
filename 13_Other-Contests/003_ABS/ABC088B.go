package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	alice, bob := 0, 0
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			bob += a[i]
		} else {
			alice += a[i]
		}
	}
	fmt.Println(alice - bob)
}
