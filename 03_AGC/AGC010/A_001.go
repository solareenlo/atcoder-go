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

	sum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		sum += a
	}

	if sum%2 != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
