package main

import (
	"bufio"
	"fmt"
	"os"
)

func C(x, y int) bool {
	s := make(map[int]bool)
	for ; x > 0; x /= 10 {
		s[x%10] = true
	}
	for ; y > 0; y /= 10 {
		s[y%10] = true
	}
	return len(s) == 1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	ans := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for j := 1; j <= a; j++ {
			if C(i, j) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
