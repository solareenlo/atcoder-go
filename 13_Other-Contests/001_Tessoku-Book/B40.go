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
	var a [100]int
	ans := 0
	for i := 0; i < n; i++ {
		var A int
		fmt.Fscan(in, &A)
		A %= 100
		ans += a[(100-A)%100]
		a[A]++
	}
	fmt.Println(ans)
}
