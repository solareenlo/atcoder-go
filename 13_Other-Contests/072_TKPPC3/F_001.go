package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		a, b int
	}

	var n int
	fmt.Fscan(in, &n)
	f := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &f[i].a, &f[i].b)
	}
	sort.Slice(f, func(i, j int) bool {
		return f[i].a*f[j].b < f[i].b*f[j].a
	})
	w := 0
	res := 0
	for i := 0; i < n; i++ {
		res += w * f[i].b
		w += f[i].a
	}
	fmt.Println(res)
}
