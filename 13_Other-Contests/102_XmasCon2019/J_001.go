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
	for n > 0 {
		n--
		var s1, s2 string
		fmt.Fscan(in, &s1, &s2)
		i, j := 0, 0
		for i < len(s1) {
			if s2[j] == s1[i] {
				j++
			}
			if j == len(s2) {
				break
			}
			i++
		}
		if j == len(s2) {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
