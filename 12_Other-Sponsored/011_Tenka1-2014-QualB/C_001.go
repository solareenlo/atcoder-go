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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	var ans [800][800]string
	for i := 0; i < 800; i++ {
		for j := 0; j < 800; j++ {
			ans[i][j] = "."
		}
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n; j++ {
			a := 0
			if i != 0 && ans[i-1][j] == "#" {
				a ^= 1
			}
			if j != 0 && ans[i][j-1] == "#" {
				a ^= 1
			}
			if j < n-1 && ans[i][j+1] == "#" {
				a ^= 1
			}
			tmp := 0
			if s[i][j] == '#' {
				tmp = 1
			}
			if a^tmp != 0 {
				ans[i+1][j] = "#"
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(ans[i][j])
		}
		fmt.Println()
	}
}
