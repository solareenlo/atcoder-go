package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p int
	fmt.Fscan(in, &n, &p)
	var a [45]int
	sum := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		if sum > p {
			fmt.Println(i - 1)
			return
		}
	}
	fmt.Println(n)
}
