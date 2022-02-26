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

	res := 0
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		cnt := 0
		for _, c := range s {
			if c == '1' {
				cnt++
			}
		}
		if cnt%2 != 0 {
			res++
		}
	}

	fmt.Println(res * (n - res))
}
