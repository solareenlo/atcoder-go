package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscan(in, &a, &b)
	n := len(a)

	a = a + "_"
	b = b + "_"
	i, j := 0, 0
	for {
		if a[i] != b[j] {
			fmt.Println("No")
			os.Exit(0)
		}
		if i == n {
			break
		}
		x, y := i, j
		for i < len(a) && a[i] == a[x] {
			i++
		}
		for j < len(b) && b[j] == b[y] {
			j++
		}
		if i-x != j-y {
			if i-x == 1 || i-x > j-y {
				fmt.Println("No")
				os.Exit(0)
			}
		}
	}

	fmt.Println("Yes")
}
