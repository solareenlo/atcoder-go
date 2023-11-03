package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var p [200200]bool

	var n int
	fmt.Fscan(in, &n)
	cnt := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if !p[i] && !p[a] {
			p[a] = true
			cnt++
		}
	}

	fmt.Println(n - cnt)
	for i := 1; i <= n; i++ {
		if !p[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
