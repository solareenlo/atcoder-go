package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var pat [107][]string

	var h, w int
	fmt.Fscan(in, &h, &w)
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(in, &s)
		pat[i] = strings.Split(s, "")
		for j := 0; j < w-1; j++ {
			if pat[i][j] == "T" && pat[i][j+1] == "T" {
				pat[i][j] = "P"
				pat[i][j+1] = "C"
			}
		}
	}
	for i := 0; i < h; i++ {
		fmt.Println(strings.Join(pat[i], ""))
	}
}
