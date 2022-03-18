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

	last := 0
	ans := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if last != 0 && a != 0 {
			ans++
			a--
		}
		ans += a / 2
		last = a % 2
	}
	fmt.Println(ans)
}
