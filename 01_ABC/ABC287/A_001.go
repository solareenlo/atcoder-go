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
	for i := 0; i < n; i++ {
		var k string
		fmt.Fscan(in, &k)
		if k == "For" {
			ans++
		}
	}
	if ans > n/2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
