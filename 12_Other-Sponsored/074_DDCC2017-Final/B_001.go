package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, z int
	fmt.Fscan(in, &n, &z)
	ans := 1
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		nw := gcd(a, z)
		ans = ans / (gcd(ans, nw)) * nw
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
