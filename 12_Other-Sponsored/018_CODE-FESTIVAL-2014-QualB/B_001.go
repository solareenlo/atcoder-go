package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a [100001]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		if ans >= m {
			fmt.Println(i)
			break
		} else {
			ans += a[i]
		}
	}
}
