package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	ans := 0
	for i := 0; i < n; {
		p := i
		for i < n && a[i] == a[p] {
			i++
		}
		ans += (i - p) * (i - p - 1) / 2
	}
	fmt.Println(ans)
}
