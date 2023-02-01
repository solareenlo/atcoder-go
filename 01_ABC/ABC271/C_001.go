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

	var t [300005]int
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a <= n {
			t[a] = 1
		}
	}

	for i := 1; i <= n; i++ {
		if t[i] == 0 {
			n--
		}
	}

	fmt.Println(n)
}
