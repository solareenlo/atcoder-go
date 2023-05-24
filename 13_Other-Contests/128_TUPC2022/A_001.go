package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	var p [200009]int
	m := 0
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &p[i])
		if p[i] != i {
			m = i
		}
	}
	if m+1 <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
