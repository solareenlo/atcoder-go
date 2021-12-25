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

	cnt := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a%2 != 0 {
			cnt++
		}
	}
	if cnt != 0 {
		fmt.Println("first")
	} else {
		fmt.Println("second")
	}
}
