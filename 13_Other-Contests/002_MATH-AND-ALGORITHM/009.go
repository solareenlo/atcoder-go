package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s int
	fmt.Fscan(in, &n, &s)

	d := make([]int, 10010)
	d[0] = 1
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for j := s - a; j >= 0; j-- {
			d[j+a] |= d[j]
		}
	}

	if d[s] != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
