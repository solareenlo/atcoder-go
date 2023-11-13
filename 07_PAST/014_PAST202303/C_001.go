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

	mp := make(map[int]int)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		mp[x] = i
	}

	for i := 1; i <= n; i++ {
		if i != 1 {
			fmt.Printf(" ")
		}
		fmt.Print(mp[i])
	}
	fmt.Println()
}
