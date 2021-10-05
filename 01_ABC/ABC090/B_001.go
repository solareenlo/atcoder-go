package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	cnt := 0
	for i := a; i <= b; i++ {
		str := strconv.Itoa(i)
		rev := reverse(strconv.Itoa(i))
		if str == rev {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func reverse(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
