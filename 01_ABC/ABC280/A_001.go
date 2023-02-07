package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var rows, columns int
	fmt.Fscan(in, &rows, &columns)

	s := make([]string, rows)
	cnt := 0
	for i := 0; i < rows; i++ {
		fmt.Fscan(in, &s[i])
		for j := 0; j < columns; j++ {
			if s[i][j] == '#' {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
