package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s [10]string
	for i := 1; i < 10; i++ {
		var t string
		fmt.Fscan(in, &t)
		s[i] = " " + t
	}

	ans := 0
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := j + 1; k+i <= 9; k++ {
				for l := 1; l+i+j <= 9; l++ {
					if s[k][l] == '#' && s[k+i][l+j] == '#' && s[k+i-j][l+i+j] == '#' && s[k-j][l+i] == '#' {
						ans++
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
