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

	var a [105]int
	var p [105][105]int
	ans := 0
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(in, &x)
		for j := 1; j <= x; j++ {
			fmt.Fscan(in, &a[j])
			for k := 1; k < j; k++ {
				if p[a[k]][a[j]] == 0 {
					ans++
				}
				p[a[k]][a[j]] = 1
			}
		}
	}

	if ans == (n-1)*n/2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
