package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n int
	fmt.Fscan(in, &n)
	a := 1
	b := 1
	for i := 3; i <= n; i++ {
		c := a
		a += b
		b = c
		a %= MOD
	}
	fmt.Println(a)
}
