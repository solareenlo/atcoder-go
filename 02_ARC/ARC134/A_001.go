package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, h, w int
	fmt.Fscan(in, &n, &h, &w)

	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	n++
	a[n] = h
	ans := (a[1] + w - 1) / w
	for i := 2; i <= n; i++ {
		ans += (a[i] - a[i-1] - 1) / w
	}
	fmt.Println(ans)
}
