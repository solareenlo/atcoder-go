package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	mp := make(map[int]bool)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		mp[a] = true
	}

	st := make(map[int]bool)
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		if mp[a] {
			st[a] = true
		}
	}

	keys := make([]int, 0, len(st))
	for k := range st {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := range keys {
		fmt.Print(keys[i], " ")
	}
}
