package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, T int
	fmt.Fscan(in, &n, &m, &T)

	var a [100005]int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var t [100005]int
	for m > 0 {
		var id, l int
		fmt.Fscan(in, &id, &l)
		t[id] = l
		m--
	}

	for i := 1; i <= n; i++ {
		if T <= 0 {
			fmt.Println("No")
			return
		}
		T += t[i] - a[i]
	}
	fmt.Println("Yes")
}
