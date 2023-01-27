package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Scan(&n)

	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	f := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 'W' && a[j][i] != 'L' {
				f |= 1
			} else {
				f |= 0
			}
			if a[i][j] == 'L' && a[j][i] != 'W' {
				f |= 1
			} else {
				f |= 0
			}
			if a[i][j] == 'D' && a[j][i] != 'D' {
				f |= 1
			} else {
				f |= 0
			}
		}
	}

	if f != 0 {
		fmt.Println("incorrect")
	} else {
		fmt.Println("correct")
	}
}
