package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	h := make([]int, n)
	m := make(map[int]int)
	for i := range h {
		fmt.Fscan(in, &h[i])
		m[h[i]] = 0
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	ord := 0
	for _, k := range keys {
		m[k] = ord
		ord++
	}
	for _, i := range h {
		fmt.Fprintln(out, m[i]+1)
	}
}
