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

	q := make(map[int]int)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		q[a]++
	}

	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		if q[a] == 0 {
			fmt.Println("No")
			return
		} else {
			q[a]--
		}
	}

	fmt.Println("Yes")
}
