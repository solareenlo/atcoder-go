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
	var L, R int
	fmt.Fscan(in, &L, &R)
	ans := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		x %= (L + R)
		ans ^= x / L
	}
	if ans == 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
