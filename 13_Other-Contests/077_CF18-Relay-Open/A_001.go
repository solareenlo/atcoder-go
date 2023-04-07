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

	day := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		var j int
		for j = 0; j < 7; j++ {
			if day[j] == s {
				break
			}
		}
		fmt.Println(day[(j+1)%7])
	}
}
