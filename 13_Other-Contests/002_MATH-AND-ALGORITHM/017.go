package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, res int
	fmt.Fscan(in, &n, &res)

	for i := 0; i < n-1; i++ {
		var a int
		fmt.Fscan(in, &a)
		res = lcm(res, a)
	}

	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
