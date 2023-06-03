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
	ans := 0
	for n > 0 {
		n--
		var a int
		fmt.Fscan(in, &a)
		ans ^= a
	}
	if ans > 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
