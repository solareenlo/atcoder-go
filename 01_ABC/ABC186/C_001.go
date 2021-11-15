package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	for i := 1; i < n+1; i++ {
		if !strings.Contains(fmt.Sprintf("%o%d", i, i), "7") {
			cnt++
		}
	}

	fmt.Println(cnt)
}
