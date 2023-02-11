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

	p := make([]int, n)
	for i := range p {
		fmt.Fscan(in, &p[i])
	}
	sort.Ints(p)

	z := 0
	for i := 1; i < n; i++ {
		z += (p[i] - p[i-1]) * i * (n - i)
	}
	fmt.Println(z)
}
