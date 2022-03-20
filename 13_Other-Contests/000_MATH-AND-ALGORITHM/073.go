package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	const mod = 1_000_000_007
	ans := 0
	x := 1
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		ans += a * x % mod
		ans %= mod
		x *= 2
		x %= mod
	}
	fmt.Println(ans)
}
