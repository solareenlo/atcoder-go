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

	var a [500001]int
	tot, cnt := 0, 0
	for i := 1; i <= n; i++ {
		var w int
		fmt.Fscan(in, &w)
		if w == i {
			cnt += tot
			tot++
		} else if a[w] == i {
			cnt++
		}
		a[i] = w
	}
	fmt.Println(cnt)
}
