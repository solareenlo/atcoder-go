package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	mp := map[[2]int]bool{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var b int
			fmt.Fscan(in, &b)
			h := (b-1)/7 - i
			r := (b-1)%7 - j
			mp[[2]int{h, r}] = true
		}
	}
	if len(mp) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
