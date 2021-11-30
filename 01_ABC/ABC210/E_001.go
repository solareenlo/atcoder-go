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

	type pair struct{ x, y int }
	v := make([]pair, m)

	for i := range v {
		fmt.Fscan(in, &v[i].y, &v[i].x)
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].x < v[j].x
	})

	res := 0
	for i := range v {
		g := gcd(n, v[i].y)
		res += (n - g) * v[i].x
		n = g
	}
	if n == 1 {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
