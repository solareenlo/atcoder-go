package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	for i := 1; i < 1000000; i++ {
		s := strconv.Itoa(i)
		s += s
		num, _ := strconv.Atoi(s)
		if num <= n {
			cnt++
		}
	}

	fmt.Println(cnt)
}
