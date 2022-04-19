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
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[x-1]++
	}

	for i := 0; i < n; i++ {
		if a[i] == 2 {
			fmt.Print(i+1, " ")
		}
	}
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			fmt.Print(i + 1)
			return
		}
	}
	fmt.Println("Correct")
}
