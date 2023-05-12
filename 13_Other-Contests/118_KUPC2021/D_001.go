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
	res := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a %= b + 1
		if (b&1) == 0 && a == b {
			res ^= 2
		} else {
			res ^= a & 1
		}
	}
	if res != 0 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Bob")
	}
}
