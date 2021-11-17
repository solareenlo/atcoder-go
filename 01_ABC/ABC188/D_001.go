package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, C int
	fmt.Fscan(in, &n, &C)

	event := map[int]int{}
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		event[a] += c
		event[b+1] -= c
	}

	keys := make([]int, 0, len(event))
	for k := range event {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	res, fee, t := 0, 0, 0
	for i := 0; i < len(event); i++ {
		if keys[i] != t {
			res += min(C, fee) * (keys[i] - t)
			t = keys[i]
		}
		fee += event[keys[i]]
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
