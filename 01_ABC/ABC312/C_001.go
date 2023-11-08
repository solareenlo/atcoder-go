package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)
	x := make([]int, a+b+1)
	for i := 1; i <= a+b; i++ {
		fmt.Fscan(in, &x[i])
		if i > a {
			x[i]++
		}
	}
	sort.Ints(x[1:])
	fmt.Println(x[b])
}
