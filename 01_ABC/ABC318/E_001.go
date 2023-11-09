package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 300005

	var T, B, C [N]int

	var n int
	fmt.Fscan(in, &n)
	res := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		T[a] += (i - B[a] - 1) * C[a]
		B[a] = i
		res += T[a]
		C[a]++
	}
	fmt.Println(res)
}
